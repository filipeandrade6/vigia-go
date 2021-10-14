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
)

type Processador struct {
	servidorGravacaoID string
	armazenamento      string
	horasRetencao      int

	registroCore registro.Core
	errChan      chan error
	matchChan    chan string

	mu        *sync.RWMutex
	processos map[string]*Processo
	retry     map[string]*Processo
	matchlist map[string]bool

	interErrChan chan traffic.ProcessoError
	regChan      chan registro.Registro
}

func New(
	servidorGravacaoID string,
	armazenamento string,
	horasRetencao int,
	registroCore registro.Core,
	errChan chan error,
	matchChan chan string,
) *Processador {
	return &Processador{
		servidorGravacaoID: servidorGravacaoID,
		armazenamento:      armazenamento,
		horasRetencao:      horasRetencao,

		registroCore: registroCore,
		errChan:      errChan,
		matchChan:    matchChan,

		mu:        &sync.RWMutex{},
		processos: make(map[string]*Processo),
		retry:     make(map[string]*Processo),
		matchlist: make(map[string]bool),

		interErrChan: make(chan traffic.ProcessoError),
		regChan:      make(chan registro.Registro),
	}
}

// =================================================================================
// Processador

func (p *Processador) Start() {
	tickerHK := time.NewTicker(time.Hour)
	tickerRetry := time.NewTicker(30 * time.Second)

	for {
		select {
		case reg := <-p.regChan:
			go p.createAndCheckRegistro(reg)

		// TODO erros da dahua estão duplicando a cada chamada
		case err := <-p.interErrChan:
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
			go p.begintHousekeeper()

		case <-tickerRetry.C:
			for processoID, processo := range p.retry {
				p.mu.Lock()
				p.processos[processoID] = processo
				p.mu.Unlock()

				processo.Start()
			}
		}
	}
}

func (p *Processador) Stop() error {
	var prc []string
	p.mu.RLock()
	for k := range p.processos {
		prc = append(prc, k)
	}
	for k := range p.retry {
		prc = append(prc, k)
	}
	p.mu.RUnlock()

	err := p.StopProcessos(prc)
	if err != nil {
		return err
	}

	return nil
}

// =================================================================================
// Processo

func (p *Processador) StartProcessos(pReq []Processo) {
	for _, prc := range pReq {
		p.mu.RLock()
		_, ok := p.processos[prc.ProcessoID]
		_, ok2 := p.retry[prc.ProcessoID]
		p.mu.RUnlock()
		if ok || ok2 {
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
			p.interErrChan,
		)

		p.mu.Lock()
		p.processos[prc.ProcessoID] = np
		p.mu.Unlock()

		// TODO fazer teste de login antes de iniciar

		np.Start()
	}
}

func (p *Processador) StopProcessos(processos []string) error {
	for _, prc := range processos {
		p.mu.RLock()
		_, ok := p.processos[prc]
		_, ok2 := p.retry[prc]
		p.mu.RUnlock()

		if ok {
			p.processos[prc].Stop()

			p.mu.Lock()
			delete(p.processos, prc)
			p.mu.Unlock()
			continue
		}

		if ok2 {
			p.mu.Lock()
			delete(p.retry, prc)
			p.mu.Unlock()
			continue
		}

		return fmt.Errorf("processo processoID[%s]: %w", prc, ErrProcessoNotFound)
	}

	return nil
}

func (p *Processador) ListProcessos() ([]string, []string) {
	var prc []string
	for k := range p.processos {
		prc = append(prc, k)
	}

	var retryPrc []string
	for k := range p.retry {
		retryPrc = append(retryPrc, k)
	}

	return prc, retryPrc
}

// =================================================================================
// Matchlist

func (p *Processador) UpdateMatchlist(placas []string) {
	p.mu.Lock()
	p.matchlist = make(map[string]bool)
	for _, placa := range placas {
		p.matchlist[placa] = true
	}
	p.mu.Unlock()
}

// =================================================================================
// Armazenamento

func (p *Processador) UpdateArmazenamento(armazenamento string, horasRetencao int) error {
	prcsBkp := make(map[string]*Processo)
	p.mu.RLock()
	for k, v := range p.processos {
		prcsBkp[k] = v
	}

	var prcs []string
	for k := range prcsBkp {
		prcs = append(prcs, k)
	}
	p.mu.RUnlock()

	err := p.StopProcessos(prcs)
	if err != nil {
		return err // TODO Tratar
	}

	p.mu.Lock()
	p.armazenamento = armazenamento
	p.horasRetencao = horasRetencao
	p.mu.Unlock()

	err = os.MkdirAll(p.armazenamento, os.ModePerm)
	if err != nil {
		return err // TODO arrumar isso aqui
	}

	var nPrcs []Processo
	for _, p := range prcsBkp {
		nPrcs = append(nPrcs, Processo{
			ProcessoID:  p.ProcessoID,
			EnderecoIP:  p.EnderecoIP,
			Porta:       p.Porta,
			Canal:       p.Canal,
			Usuario:     p.Usuario,
			Senha:       p.Senha,
			Processador: p.Processador,
		})
	}

	p.StartProcessos(nPrcs)

	return nil
}

// =================================================================================

func (p *Processador) begintHousekeeper() {
	d := time.Now().Add(time.Duration(-p.horasRetencao) * time.Hour)

	err := filepath.Walk(p.armazenamento, func(path string, info os.FileInfo, err error) error {
		if path == p.armazenamento {
			return nil
		}

		if info.ModTime().Before(d) {
			err := os.Remove(path)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		p.errChan <- fmt.Errorf("housekeeper stopped: %w", err)
	}
}

func (p *Processador) createAndCheckRegistro(reg registro.Registro) {
	_, err := p.registroCore.Create(context.Background(), reg)
	if err != nil {
		fmt.Println(err, reg.ProcessoID)
		p.errChan <- err // TODO adicionar error personalizado
		return
	}

	p.mu.RLock()
	_, ok := p.matchlist[reg.Placa]
	p.mu.RUnlock()
	if ok {
		p.matchChan <- reg.RegistroID
	}
}
