package processador

import (
	"context"
	"fmt"

	"github.com/filipeandrade6/vigia-go/internal/core/camera"
	"github.com/filipeandrade6/vigia-go/internal/core/processo"
	"github.com/filipeandrade6/vigia-go/internal/core/registro"
	"github.com/filipeandrade6/vigia-go/internal/core/veiculo"
)

type Processador struct {
	ProcessoCore       processo.Core // TODO ver se é usado pointer nos core...
	CameraCore         camera.Core
	RegistroCore       registro.Core
	VeiculoCore        veiculo.Core
	ServidorGravacaoID string
	Armazenamento      string
	SuperrChan         chan error
	MatchChan          chan string

	processos map[string]*Processo
	matchlist map[string]bool

	errChan chan error
	regChan chan registro.Registro
}

func NewProcessador(servidorGravacaoID, armazenamento string, processoCore processo.Core, cameraCore camera.Core, registroCore registro.Core, veiculoCore veiculo.Core, SuperrChan chan error, MatchChan chan string) *Processador {
	return &Processador{
		ServidorGravacaoID: servidorGravacaoID,
		ProcessoCore:       processoCore,
		CameraCore:         cameraCore,
		RegistroCore:       registroCore,
		VeiculoCore:        veiculoCore,
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
			go p.Salvar(reg, p.SuperrChan)

			if _, ok := p.matchlist[reg.Placa]; ok {
				p.MatchChan <- reg.RegistroID
			}
		case err := <-p.errChan:
			p.SuperrChan <- err // TODO ve o tipo de problema e tem como recuperar - manda ou para notificação ou para SuperrChan
		}
	}
}

func (p *Processador) StartProcesso(ctx context.Context, processoID string) error {
	prclist, ok := p.processos[processoID]
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
		return fmt.Errorf("this processo don't belong in this servidor de gravacao")
	}

	cam, err := p.CameraCore.QueryByID(ctx, prc.CameraID)
	if err != nil {
		return fmt.Errorf("query camera processoID[%s]: %w", processoID, err)
	}

	np := NewProcesso(prc, cam, p.ServidorGravacaoID, p.Armazenamento, p.regChan, p.errChan)

	if err := np.Start(); err != nil {
		return fmt.Errorf("initializing processo processoID[%q]: %w", processoID, err)
	}

	p.processos[processoID] = np

	return nil
}

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

func (p *Processador) AtualizarMatchList(ctx context.Context) error {
	veiculos, err := p.VeiculoCore.Query(ctx, "", 1, 1000000) // TODO arrumar depois
	if err != nil {
		return fmt.Errorf("querying veiculos database")
	}

	for _, veiculo := range veiculos {
		p.matchlist[veiculo.Placa] = true
	}

	return nil
}
