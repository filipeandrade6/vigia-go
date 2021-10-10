package service

import (
	"fmt"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/core/registro"
	"github.com/filipeandrade6/vigia-go/internal/sys/validate"
)

type Processo struct {
	ProcessoID    string
	EnderecoIP    string
	Porta         int
	Canal         int
	Usuario       string
	Senha         string
	Armazenamento string
	Processador   int
	regChan       chan registro.Registro
	errChan       chan error
	stopChan      chan struct{}
	stoppedChan   chan struct{}
}

func NewProcesso(
	processoID string,
	enderecoIP string,
	porta int,
	canal int,
	usuario string,
	senha string,
	armazenamento string,
	processador int,
	regChan chan registro.Registro,
	errChan chan error,
) *Processo {
	return &Processo{
		ProcessoID:    processoID,
		EnderecoIP:    enderecoIP,
		Porta:         porta,
		Canal:         canal,
		Usuario:       usuario,
		Senha:         senha,
		Armazenamento: armazenamento,
		Processador:   processador,
		regChan:       regChan,
		errChan:       errChan,
		stopChan:      make(chan struct{}),
		stoppedChan:   make(chan struct{}),
	}
}

func (p *Processo) Start() {
	go p.processar()
}

func (p *Processo) Stop() {
	close(p.stopChan)
	<-p.stoppedChan
}

func (p *Processo) processar() {
	stopProcChan := make(chan struct{})
	stoppedProcChan := make(chan struct{})

	if p.Processador == 0 {
		go dahua(
			p.ProcessoID,
			p.EnderecoIP,
			p.Porta,
			p.Canal,
			p.Usuario,
			p.Senha,
			p.Armazenamento,
			p.regChan,
			p.errChan,
			stopProcChan,
			stoppedProcChan,
		)
	}

	<-p.stopChan
	close(stopProcChan)
	<-stoppedProcChan
	close(p.stoppedChan)
}

func dahua(
	processoID string,
	enderecoIP string,
	porta int,
	canal int,
	usuario string,
	senha string,
	armazenamento string,
	regChan chan registro.Registro,
	errChan chan error,
	stopProcChan chan struct{},
	stoppedProcChan chan struct{},
) {
	defer close(stoppedProcChan)

	var i int
	for {
		select {
		default:
			time.Sleep(time.Duration(time.Millisecond * 200))
			r := registro.Registro{
				RegistroID:    validate.GenerateID(),
				ProcessoID:    processoID,
				Placa:         fmt.Sprintf("ABC%04d", i),
				TipoVeiculo:   "sedan",
				CorVeiculo:    "prata",
				MarcaVeiculo:  "honda",
				Armazenamento: "",
				Confianca:     0.50,
				CriadoEm:      time.Now(),
			}
			r.Armazenamento = fmt.Sprintf("%s/%d_%s", armazenamento, r.CriadoEm.Unix(), r.RegistroID)
			regChan <- r
			i++
		case <-stopProcChan:
			return
		}
	}
}