package processador

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/core/registro"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/dahua/v1/traffic"
)

var (
	ErrProcessoNotFound = errors.New("processo not found")
	ErrAlreadyExecuting = errors.New("processo already executing")
)

type Processador struct {
	registroCore       registro.Core
	servidorGravacaoID string
	errChan            chan error
	matchChan          chan string

	mu            *sync.RWMutex
	processos     map[string]*Processo
	retry         map[string]*Processo // ! implementar
	matchlist     map[string]bool
	armazenamento string

	housekeeperStatus bool
	horasRetencao     int

	internalErrChan chan traffic.ProcessoError // melhorar esse erro
	regChan         chan registro.Registro
}

func New(
	registroCore registro.Core,
	servidorGravacaoID string,
	armazenamento string,
	horasRetencao int,
	errChan chan error,
	matchChan chan string,
) *Processador {
	return &Processador{
		registroCore:       registroCore,
		servidorGravacaoID: servidorGravacaoID,
		armazenamento:      armazenamento,
		errChan:            errChan,
		matchChan:          matchChan,

		mu:        &sync.RWMutex{},
		processos: make(map[string]*Processo),
		retry:     make(map[string]*Processo), // ! implementar
		matchlist: make(map[string]bool),

		housekeeperStatus: true,
		horasRetencao:     horasRetencao,

		regChan:         make(chan registro.Registro),
		internalErrChan: make(chan traffic.ProcessoError),
	}
}

// TODO como vai funcionar o Back-off? https://github.com/jpillora/backoff/blob/master/backoff.go
// =================================================================================

func (p *Processador) Start() {
	tickerHK := time.NewTicker(time.Hour)
	tickerRetry := time.NewTicker(15 * time.Second)

	for {
		select {
		case reg := <-p.regChan:
			go p.createAndCheckRegistro(reg)

		case err := <-p.internalErrChan:
			switch {
			case errors.As(err, &traffic.ErrLogin):
				p.retry[err.ProcessoID] = p.processos[err.ProcessoID]
				delete(p.processos, err.ProcessoID)

			case errors.As(err, &traffic.ErrSaveImage):
				// TODO notificar

			case errors.As(err, &traffic.ErrAnalyzer):
				delete(p.processos, err.ProcessoID)
				// TODO notificar
			}
			p.errChan <- err

		case <-tickerHK.C:
			if p.housekeeperStatus {
				go p.begintHousekeeper()
			}

		case <-tickerRetry.C:
			fmt.Printf("\n retry: %v\n", p.retry)
			for processoID, processo := range p.retry {
				p.mu.Lock()
				p.processos[processoID] = processo
				p.mu.Unlock()

				processo.Start()
			}
		}
	}
}

// TODO remover do retry também...
func (p *Processador) Stop() error {
	var prc []string
	for k := range p.processos {
		prc = append(prc, k)
	}

	err := p.StopProcessos(prc)
	if err != nil {
		return err
	}

	return nil
}

// =================================================================================

func (p *Processador) StartProcessos(pReq []Processo) {
	for _, prc := range pReq {
		p.mu.RLock()
		_, ok := p.processos[prc.ProcessoID]
		p.mu.RUnlock()
		if ok {
			continue
		}

		np := NewProcesso(
			prc.ProcessoID,
			prc.EnderecoIP,
			prc.Porta,
			prc.Canal,
			prc.Usuario,
			prc.Senha,
			prc.Processador,

			p.armazenamento,
			p.regChan,
			p.internalErrChan,
		)

		p.mu.Lock()
		p.processos[prc.ProcessoID] = np
		p.mu.Unlock()

		np.Start()
	}
}

func (p *Processador) StopProcessos(processos []string) error {
	for _, prc := range processos {
		p.mu.RLock()
		_, ok := p.processos[prc]
		p.mu.RUnlock()
		if !ok {
			return fmt.Errorf("processo processoID[%s]: %w", prc, ErrProcessoNotFound)
		}

		p.processos[prc].Stop()

		p.mu.Lock()
		delete(p.processos, prc)
		p.mu.Unlock()
	}

	return nil
}

func (p *Processador) ListProcessos() []string {
	var prc []string
	for k := range p.processos {
		prc = append(prc, k)
	}

	return prc
}

// =================================================================================

func (p *Processador) AtualizarMatchList(placas []string) error {
	p.mu.Lock()
	p.matchlist = make(map[string]bool)
	for _, placa := range placas {
		p.matchlist[placa] = true
	}
	p.mu.Unlock()

	return nil
}

// =================================================================================

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

// TODO colocar mais inforamções -  mudar para processador info - servidor info deve ficar no server
func (p *Processador) GetServidorInfo() (string, int) {
	return p.armazenamento, p.horasRetencao
}

// =================================================================================

func (p *Processador) createAndCheckRegistro(reg registro.Registro) {
	_, err := p.registroCore.Create(context.Background(), reg)
	if err != nil {
		fmt.Println(err, reg.ProcessoID)
		p.errChan <- err // internal ou errChan?
		return
	}

	p.mu.RLock()
	_, ok := p.matchlist[reg.Placa]
	p.mu.RUnlock()
	if ok {
		p.matchChan <- reg.RegistroID
	}
}
