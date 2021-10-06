package processador

import (
	"context"
	"fmt"

	"github.com/filipeandrade6/vigia-go/internal/core/camera"
	"github.com/filipeandrade6/vigia-go/internal/core/processo"
	"github.com/filipeandrade6/vigia-go/internal/core/registro"
)

type Processador struct {
	ProcessoID         string
	ServidorGravacaoID string
	ProcessoCore       processo.Core
	CameraCore         camera.Core
	RegistroCore       registro.Core // TODO pointer?
	Armazenamento      string
	SuperrChan         chan error

	processos map[string]*Processo // TODO pointer e todo o resto?
	matchlist map[string]bool

	errChan chan error
	regChan chan registro.Registro
}

// TODO como o adaptador vai funcionar...
func NewProcessador(servidorGravacaoID, armazenamento string, processoCore processo.Core, cameraCore camera.Core, registroCore registro.Core, SuperrChan chan error) *Processador {
	return &Processador{
		ServidorGravacaoID: servidorGravacaoID, // é usado aonde?
		ProcessoCore:       processoCore,
		CameraCore:         cameraCore,
		RegistroCore:       registroCore,
		Armazenamento:      armazenamento,
		SuperrChan:         SuperrChan,

		processos: make(map[string]*Processo),
		matchlist: make(map[string]bool),
		regChan:   make(chan registro.Registro),
		errChan:   make(chan error),
	}
}

func (p *Processador) Gerenciar() {
	for {
		select {
		case reg := <-p.regChan:
			p.Salvar(reg, p.SuperrChan)
		case err := <-p.errChan:
			// TODO ve o tipo de problema e tem como recuperar - manda ou para notificação ou para SuperrChan
			p.SuperrChan <- err
		}
	}
}

// func (p *Processador) NovoProcesso(ctx context.Context, processoID string) error {
// }

func (p *Processador) StartProcesso(ctx context.Context, processoID string) error {
	prclist, ok := p.processos[processoID] // se já estiver na lista de processos inicia o processo.
	if ok {
		if err := prclist.Start(); err != nil {
			return fmt.Errorf("processo processoID [%q] already added but failed to start: %w", processoID, err)
		}
	}

	prc, err := p.ProcessoCore.QueryByID(ctx, processoID)
	if err != nil {
		return fmt.Errorf("query processo processoID[%s]: %w", processoID, err)
	}

	cam, err := p.CameraCore.QueryByID(ctx, prc.CameraID)
	if err != nil {
		return fmt.Errorf("query camera processoID[%s]: %w", processoID, err)
	}

	np := NewProcesso(
		prc.ProcessoID,
		p.ServidorGravacaoID,
		p.Armazenamento,
		prc.Processador,
		prc.Adaptador,
		cam,
		p.regChan,
		p.errChan,
	)

	if err := np.Start(); err != nil {
		return fmt.Errorf("initializing processo processoID[%q]: %w", processoID, err)
	}

	return nil
}

// TODO receber contexto...
func (p *Processador) StopProcesso(processoID string) error {
	prclist, ok := p.processos[processoID] // se já estiver na lista de processos pausa ou informa erro.
	if ok {
		if !prclist.Status {
			return fmt.Errorf("processo processoID[%q] already stopped", processoID)
		}

		if err := prclist.Stop(); err != nil {
			return fmt.Errorf("stoping processo processoID[%q]: %w", processoID, err)
		}

		return nil
	}

	return fmt.Errorf("processo processoID[%q] not found in servidor de gravacao", processoID)
}

func (p *Processador) Salvar(reg registro.Registro, errChan chan error) {
	_, err := p.RegistroCore.Create(context.Background(), reg) // TODO alterar no banco de dados - quando criar não gerar
	if err != nil {
		errChan <- err
	}
}
