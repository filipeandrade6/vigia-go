package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/core/registro"
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
	errChan chan error,
	matchChan chan string,
	stopChan chan struct{}, // TODO utilizar?
	stoppedChan chan struct{},
) *Processador {
	return &Processador{
		servidorGravacaoID: servidorGravacaoID,
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
			go p.createAndCheckRegistro(reg)

		case err := <-p.internalErrChan:
			p.errChan <- err // TODO ve o tipo de problema e tem como recuperar - manda ou para notificação ou para SuperrChan

		case <-ticker.C:
			if p.housekeeperStatus {
				go p.begintHousekeeper()
			}
		}
	}
}

func (p *Processador) StartProcessos(pReq []Processo) error {
	for _, prc := range pReq {
		p.mutex.RLock()
		_, ok := p.processos[prc.ProcessoID]
		p.mutex.RUnlock()
		if ok {
			return fmt.Errorf("processo processoID[%s]: %s", prc.ProcessoID, ErrAlreadyExecuting)
		}

		np := NewProcesso(
			prc.ProcessoID,
			prc.EnderecoIP,
			prc.Porta,
			prc.Canal,
			prc.Usuario,
			prc.Senha,

			p.armazenamento,
			p.regChan,
			p.errChan,
		)

		np.Start()

		p.mutex.Lock()
		p.processos[prc.ProcessoID] = np
		p.mutex.Unlock()
	}

	return nil
}

func (p *Processador) StopProcessos(processos []string) error {
	for _, prc := range processos {
		p.mutex.RLock()
		_, ok := p.processos[prc]
		p.mutex.RUnlock()
		if !ok {
			return fmt.Errorf("processo processoID[%s]: %w", prc, ErrProcessoNotFound)
		}

		p.processos[prc].Stop()
	}

	return nil
}

func (p *Processador) ListProcessos() []string {
	prc := make([]string, len(p.processos))
	for k := range p.processos {
		prc = append(prc, k)
	}

	return prc
}

// =================================================================================

func (p *Processador) AtualizarMatchList(placas []string) error {
	p.mutex.Lock()
	p.matchlist = make(map[string]bool)
	for _, placa := range placas {
		p.matchlist[placa] = true
	}
	p.mutex.Unlock()

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

func (p *Processador) GetServidorInfo() (string, int) {
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

// func (p *Processador) StopGerencia() {
// 	err := p.RemoveAllProcessos()
// 	if err != nil {

// 	}

// }
