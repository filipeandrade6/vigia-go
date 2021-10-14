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
	retry     map[string]*Processo // ! implementar
	matchlist map[string]bool

	internalErrChan chan traffic.ProcessoError // melhorar esse erro
	regChan         chan registro.Registro
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

		internalErrChan: make(chan traffic.ProcessoError),
		regChan:         make(chan registro.Registro),
	}
}

// =================================================================================

func (p *Processador) Start() {
	err := os.MkdirAll(p.armazenamento, os.ModePerm)
	if err != nil {
		p.errChan <- err
		return // TODO arrumar isso aqui
	}

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

// TODO beginHousekeeper deve receber contexto para parar em caso de alguma coisa...?
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
