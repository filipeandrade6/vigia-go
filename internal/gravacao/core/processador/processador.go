package processador

import (
	"context"
	"errors"
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
	MatchChan          chan string

	processos map[string]*Processo // TODO pointer e todo o resto?
	matchlist map[string]bool

	errChan chan error
	regChan chan registro.Registro
}

// TODO como o adaptador vai funcionar?
func NewProcessador(servidorGravacaoID, armazenamento string, processoCore processo.Core, cameraCore camera.Core, registroCore registro.Core, SuperrChan chan error, MatchChan chan string) *Processador {
	return &Processador{
		ServidorGravacaoID: servidorGravacaoID, // é usado aonde?
		ProcessoCore:       processoCore,
		CameraCore:         cameraCore,
		RegistroCore:       registroCore,
		Armazenamento:      armazenamento,
		SuperrChan:         SuperrChan,
		MatchChan:          MatchChan,

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
			if _, ok := p.matchlist[reg.Placa]; ok {
				p.MatchChan <- reg.RegistroID
			}

			go p.Salvar(reg, p.SuperrChan) // coloquei em uma nova goroutine, ve se aumenta a performance

		case err := <-p.errChan:
			// TODO ve o tipo de problema e tem como recuperar - manda ou para notificação ou para SuperrChan
			p.SuperrChan <- err
		}
	}
}

func (p *Processador) StartProcesso(ctx context.Context, processoID string) error {
	prclist, ok := p.processos[processoID] // se já estiver na lista de processos inicia o processo.
	if ok {
		if err := prclist.Start(); err != nil {
			if err.Error() == "processo already executing" {
				return fmt.Errorf("processo [%q] already executing", prclist.ProcessoID)
			}
			return fmt.Errorf("processo processoID [%q] already added but failed to start: %w", processoID, err)
		}
	}

	prc, err := p.ProcessoCore.QueryByID(ctx, processoID)
	if err != nil {
		return fmt.Errorf("query processo processoID[%s]: %w", processoID, err)
	}

	if prc.ServidorGravacaoID != p.ServidorGravacaoID {
		return errors.New("this processo don't belong in this servidor de gravacao")
	}

	cam, err := p.CameraCore.QueryByID(ctx, prc.CameraID)
	if err != nil {
		return fmt.Errorf("query camera processoID[%s]: %w", processoID, err)
	}

	np := NewProcesso(
		prc.ProcessoID,
		prc.ServidorGravacaoID,
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

	p.processos[processoID] = np

	return nil
}

// TODO receber contexto...
func (p *Processador) StopProcesso(ctx context.Context, processoID string) error {
	prclist, ok := p.processos[processoID]
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
