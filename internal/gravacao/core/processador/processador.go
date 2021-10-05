package processador

import (
	"context"
	"fmt"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/core/camera"
	"github.com/filipeandrade6/vigia-go/internal/core/registro"
)

type Processador struct {
	ServidorGravacaoID string
	CameraCore         camera.Core
	RegistroCore       registro.Core // TODO pointer?
	Armazenamento      string
	SupErrChan         chan error

	processos map[string]*Processo // TODO pointer e todo o resto?
	matchlist map[string]bool

	ErrChan chan error
	regChan chan registro.Registro
}

// TODO como o adaptador vai funcionar...
func NewProcessador(servidorGravacaoID, armazenamento string, cameraCore camera.Core, registroCore registro.Core, SupErrChan chan error) *Processador {
	return &Processador{
		ServidorGravacaoID: servidorGravacaoID,
		CameraCore:         cameraCore,
		RegistroCore:       registroCore,
		Armazenamento:      armazenamento,
		SupErrChan:         SupErrChan,

		processos: make(map[string]*Processo),
		matchlist: make(map[string]bool),
		regChan:   make(chan registro.Registro),
		ErrChan:   make(chan error),
	}
}

func (p *Processador) Gerenciar() {
	for {
		select {
		case reg := <-p.regChan:
			p.Salvar(reg, p.SupErrChan)
		case err := <-p.ErrChan:
			// TODO ve o tipo de problema e tem como recuperar - manda ou para notificação ou para SupErrChan
			p.SupErrChan <- err
		}
	}
}

func (p *Processador) NovoProcesso(cameraID string, processador, adaptador int) (string, error) {
	cam, err := p.CameraCore.QueryByID(context.Background(), cameraID)
	if err != nil {
		return "", fmt.Errorf("query cameraID[%s]: %w", cameraID, err)
	}

	prc := NewProcesso(
		p.ServidorGravacaoID,
		p.Armazenamento,
		processador,
		adaptador,
		cam,
		p.regChan,
		p.ErrChan,
	)
	p.processos[prc.ProcessoID] = prc

	return prc.ProcessoID, nil
}

func (p *Processador) StartProcesso(processoID string) {
	prc := p.processos[processoID]
	prc.Start() // TODO adicionar mais coisas...?
}

func (p *Processador) StopProcesso(processoID string) {
	prc := p.processos[processoID]
	prc.Stop() // TODO adicionar mais coisas...?
}

func (p *Processador) Salvar(reg registro.Registro, ErrChan chan error) {
	regToDB := registro.Registro{
		RegistroID:    reg.RegistroID,
		ProcessoID:    reg.ProcessoID,
		Placa:         reg.Placa,
		TipoVeiculo:   reg.TipoVeiculo,
		CorVeiculo:    reg.CorVeiculo,
		MarcaVeiculo:  reg.MarcaVeiculo,
		Armazenamento: reg.Armazenamento,
		Confianca:     0.5,
		CriadoEm:      time.Now(),
	}
	_, err := p.RegistroCore.Create(context.Background(), regToDB)
	if err != nil {
		ErrChan <- err
	}
}
