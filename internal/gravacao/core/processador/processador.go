package processador

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/core/camera"
	"github.com/filipeandrade6/vigia-go/internal/core/processo"
	"github.com/filipeandrade6/vigia-go/internal/core/registro"
	"github.com/filipeandrade6/vigia-go/internal/core/veiculo"
)

var (
	ErrProcessoNotFound = errors.New("processo not found")
	ErrAlreadyStarted   = errors.New("processo already started")
	ErrAlreadyStopped   = errors.New("processo already stopped")
)

type Processador struct {
	processoCore       processo.Core
	cameraCore         camera.Core
	registroCore       registro.Core
	veiculoCore        veiculo.Core
	servidorGravacaoID string
	errChan            chan error
	matchChan          chan string
	stopChan           chan struct{}
	stoppedChan        chan struct{}

	mutex         *sync.RWMutex
	processos     map[string]*Processo
	matchlist     map[string]bool
	armazenamento string

	housekeeperStatus bool
	horasRetencao     int

	internalErrChan chan error
	regChan         chan registro.Registro
}

func NewProcessador(
	servidorGravacaoID string,
	armazenamento string,
	horasRetencao int,
	processoCore processo.Core,
	cameraCore camera.Core,
	registroCore registro.Core,
	veiculoCore veiculo.Core,
	errChan chan error,
	matchChan chan string,
	stopChan chan struct{}, // TODO utilizar?
	stoppedChan chan struct{},
) *Processador {
	return &Processador{
		servidorGravacaoID: servidorGravacaoID,
		processoCore:       processoCore,
		cameraCore:         cameraCore,
		registroCore:       registroCore,
		veiculoCore:        veiculoCore,
		armazenamento:      armazenamento,
		errChan:            errChan,
		matchChan:          matchChan,
		stopChan:           stopChan,
		stoppedChan:        stoppedChan,

		mutex:     &sync.RWMutex{},
		processos: make(map[string]*Processo),
		matchlist: make(map[string]bool),

		housekeeperStatus: true,
		horasRetencao:     horasRetencao,

		regChan:         make(chan registro.Registro),
		internalErrChan: make(chan error),
	}
}

// TODO como vai funcionar o Back-off? https://github.com/jpillora/backoff/blob/master/backoff.go
// =================================================================================

func (p *Processador) Gerenciar() {
	ticker := time.NewTicker(time.Hour)

	for {
		select {
		case reg := <-p.regChan:
			go p.createAndCheckRegistro(reg) // TODO aqui esta meio estranho...

		case err := <-p.internalErrChan:
			p.errChan <- err // TODO ve o tipo de problema e tem como recuperar - manda ou para notificação ou para SuperrChan

		case <-ticker.C:
			if p.housekeeperStatus {
				go p.begintHousekeeper()
			}

		}
	}
}

// =================================================================================

func (p *Processador) StartProcesso(ctx context.Context, processoID string) error {
	p.mutex.RLock()
	prclist, ok := p.processos[processoID]
	p.mutex.RUnlock()
	if ok {
		if prclist.status {
			return ErrAlreadyStarted
		}
		prclist.Start()
	}

	prc, err := p.processoCore.QueryByID(ctx, processoID)
	if err != nil {
		return fmt.Errorf("query processo processoID[%s]: %w", processoID, err)
	}

	if prc.ServidorGravacaoID != p.servidorGravacaoID {
		return fmt.Errorf("this processo don't belong in this servidor de gravacao")
	}

	cam, err := p.cameraCore.QueryByID(ctx, prc.CameraID)
	if err != nil {
		return fmt.Errorf("query camera processoID[%s]: %w", processoID, err)
	}

	np := NewProcesso(prc, cam, p.servidorGravacaoID, p.armazenamento, p.regChan, p.errChan)

	np.Start()

	p.mutex.Lock()
	p.processos[processoID] = np
	p.mutex.Unlock()

	return nil
}

func (p *Processador) StartAllProcessos(ctx context.Context) error {
	prcs, err := p.processoCore.QueryAll(ctx)
	if err != nil {
		return err
	}

	for _, prc := range prcs {
		if err := p.StartProcesso(ctx, prc.ProcessoID); err != nil {
			if !errors.Is(err, ErrAlreadyStarted) {
				return err
			}
		}
	}

	return nil
}

func (p *Processador) StopProcesso(processoID string) error {
	prclist, ok := p.processos[processoID]
	if !ok {
		return ErrProcessoNotFound
	}

	if !prclist.status {
		return ErrAlreadyStopped
	}
	prclist.Stop()

	return nil
}

func (p *Processador) StopAllProcessos() error {
	for _, prc := range p.processos {
		if err := p.StopProcesso(prc.ProcessoID); !errors.Is(err, ErrAlreadyStopped) {
			return err
		}
	}

	return nil
}

func (p *Processador) RemoveProcesso(processoID string) error {
	if err := p.StopProcesso(processoID); err != nil {
		if !errors.Is(err, ErrAlreadyStopped) {
			return err
		}
	}

	delete(p.processos, processoID)

	return nil
}

func (p *Processador) RemoveAllProcessos() error {
	if err := p.StopAllProcessos(); err != nil {
		return err
	}

	for _, prc := range p.processos {
		if err := p.RemoveProcesso(prc.ProcessoID); err != nil {
			return err
		}
	}

	return nil
}

func (p *Processador) ShowAllProcessos() (map[string]bool, error) {
	prc := make(map[string]bool)

	p.mutex.RLock()
	for _, processo := range p.processos {
		prc[processo.ProcessoID] = processo.status
	}
	p.mutex.RUnlock()

	return prc, nil
}

// =================================================================================

func (p *Processador) AtualizarMatchList(ctx context.Context) error {
	veiculos, err := p.veiculoCore.QueryAll(ctx)
	if err != nil {
		return fmt.Errorf("querying veiculos database")
	}

	p.mutex.Lock()
	p.matchlist = make(map[string]bool)
	for _, veiculo := range veiculos {
		p.matchlist[veiculo.Placa] = true
	}
	p.mutex.Unlock()

	return nil
}

// =================================================================================
// TODO Update no gerencia depois chama essa funcao...

func (p *Processador) AtualizarHousekeeper(horasRetencao int) {
	p.horasRetencao = horasRetencao
}

func (p *Processador) StartHousekeeper() {
	p.housekeeperStatus = true
}

func (p *Processador) StopHousekeeper() {
	p.housekeeperStatus = false
}

func (p *Processador) StatusHousekeeper() bool {
	return p.housekeeperStatus
}

// TODO beginHousekeeper deve receber contexto para parar em caso de alguma coisa
func (p *Processador) begintHousekeeper() {
	d := time.Now().Add(time.Duration(-p.horasRetencao) * time.Hour)

	err := filepath.Walk(p.armazenamento, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // TODO testar com diretorio errado...
		}

		if info.ModTime().Before(d) {
			os.Remove(path)
		}

		return nil
	})

	if err != nil {
		p.errChan <- err
	}
}

// =================================================================================

func (p *Processador) GetArmazenamentoInfo() (string, int) {
	return p.armazenamento, p.horasRetencao
}

// =================================================================================

func (p *Processador) createAndCheckRegistro(reg registro.Registro) {
	_, err := p.registroCore.Create(context.Background(), reg)
	if err != nil {
		p.internalErrChan <- err
		return
	}

	p.mutex.RLock()
	_, ok := p.matchlist[reg.Placa]
	p.mutex.RUnlock()
	if ok {
		p.matchChan <- reg.RegistroID
	}
}

// =================================================================================

func (p *Processador) StopGerencia() {
	err := p.RemoveAllProcessos()
	if err != nil {

	}

	Parar o housekeeper


}