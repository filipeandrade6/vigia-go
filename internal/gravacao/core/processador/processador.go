package processador

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/filipeandrade6/vigia-go/internal/core/camera"
	"github.com/filipeandrade6/vigia-go/internal/core/processo"
	"github.com/filipeandrade6/vigia-go/internal/core/registro"
	"github.com/filipeandrade6/vigia-go/internal/core/veiculo"
)

var (
	ErrProcessoNotFound = errors.New("processo not found")
)

type Processador struct {
	processoCore       processo.Core // TODO ver se é usado pointer nos core...
	cameraCore         camera.Core
	registroCore       registro.Core
	veiculoCore        veiculo.Core
	servidorGravacaoID string
	armazenamento      string
	errChan            chan error
	matchChan          chan string

	mutex     *sync.RWMutex
	processos map[string]*Processo
	matchlist map[string]bool

	internalErrChan chan error
	regChan         chan registro.Registro
}

func NewProcessador(servidorGravacaoID, armazenamento string, processoCore processo.Core, cameraCore camera.Core, registroCore registro.Core, veiculoCore veiculo.Core, SuperrChan chan error, MatchChan chan string) *Processador {
	return &Processador{
		servidorGravacaoID: servidorGravacaoID,
		processoCore:       processoCore,
		cameraCore:         cameraCore,
		registroCore:       registroCore,
		veiculoCore:        veiculoCore,
		armazenamento:      armazenamento,
		errChan:            SuperrChan,
		matchChan:          MatchChan,

		mutex:           &sync.RWMutex{},
		processos:       make(map[string]*Processo),
		matchlist:       make(map[string]bool),
		regChan:         make(chan registro.Registro),
		internalErrChan: make(chan error),
	}
}

func (p *Processador) Gerenciar() {
	for {
		select {
		case reg := <-p.regChan:
			go p.createRegistro(reg, p.errChan)

			p.mutex.RLock()
			if _, ok := p.matchlist[reg.Placa]; ok { // concurrent map access
				p.matchChan <- reg.RegistroID
			}
			p.mutex.RUnlock()
		case err := <-p.internalErrChan:
			p.errChan <- err // TODO ve o tipo de problema e tem como recuperar - manda ou para notificação ou para SuperrChan
		}
	}
}

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

func (p *Processador) StopProcesso(ctx context.Context, processoID string) error {
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

func (p *Processador) StopAllProcessos(ctx context.Context) error {
	for _, prc := range p.processos {
		if err := p.StopProcesso(ctx, prc.ProcessoID); !errors.Is(err, ErrAlreadyStopped) {
			return err
		}
	}

	return nil
}

func (p *Processador) RemoveProcesso(ctx context.Context, processoID string) error {
	if err := p.StopProcesso(ctx, processoID); err != nil {
		if !errors.Is(err, ErrAlreadyStopped) {
			return err
		}
	}

	delete(p.processos, processoID)

	return nil
}

func (p *Processador) RemoveAllProcessos(ctx context.Context) error {
	for _, prc := range p.processos {
		if err := p.RemoveProcesso(ctx, prc.ProcessoID); err != nil {
			return err
		}
	}

	return nil
}

func (p *Processador) AtualizarMatchList(ctx context.Context) error {
	veiculos, err := p.veiculoCore.Query(ctx, "", 1, 1000000) // TODO arrumar depois
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

func (p *Processador) AtualizarHousekeeper() error {

}

// TODO implementar
func (p *Processador) SystemInfo(ctx context.Context) error {
	return nil
}

func (p *Processador) createRegistro(reg registro.Registro, errChan chan error) {
	_, err := p.registroCore.Create(context.Background(), reg) // TODO alterar no banco de dados - quando criar não gerar
	if err != nil {
		errChan <- err
	}
}
