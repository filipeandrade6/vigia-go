package processador

import (
	"errors"
	"fmt"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/core/camera"
	"github.com/filipeandrade6/vigia-go/internal/core/registro"
	"github.com/filipeandrade6/vigia-go/internal/sys/validate"
)

type Processo struct {
	ProcessoID         string        // identificador do processo
	ServidorGravacaoID string        // identificador do servidor de gravacao
	Camera             camera.Camera // configurações da câmera
	Processador        int           // processador utilizado
	Adaptador          int           // adaptador utilizado
	Armazenamento      string        // local onde será salvo as imagens

	Status bool // true: executando, false: parado

	regChan     chan registro.Registro // channel onde será enviado os registros TODO no Processador
	ErrChan     chan error             // channel onde será enviado erro
	stopChan    chan struct{}          // channel de sinalização para parar o processamento
	stoppedChan chan struct{}          // channel de sinalização que o processamento foi parado
}

func NewProcesso(
	processoID string,
	servidorGravacaoID string,
	armazenamento string,
	processador int,
	adaptador int,
	camera camera.Camera,
	regChan chan registro.Registro,
	ErrChan chan error,
) *Processo {
	return &Processo{
		ProcessoID:         processoID,
		ServidorGravacaoID: servidorGravacaoID,
		Armazenamento:      armazenamento,
		Processador:        processador,
		Adaptador:          adaptador,
		Camera:             camera,
		regChan:            regChan,
		ErrChan:            ErrChan,
	}
}

func (p *Processo) Start() error {
	if p.Status {
		return errors.New("processo already executing")
	}

	p.stopChan = make(chan struct{})
	p.stoppedChan = make(chan struct{})

	go p.processar()
	p.Status = true

	return nil
}

func (p *Processo) Stop() error {
	if !p.Status {
		return errors.New("processo already paused")
	}

	// TODO colocar context?
	close(p.stopChan)
	<-p.stoppedChan

	p.Status = false

	return nil
}

func (p *Processo) processar() {
	defer close(p.stoppedChan)

	ErrChan := make(chan error)
	stopCmdChan := make(chan struct{})

	go dahua(p.regChan, ErrChan, stopCmdChan, p.ProcessoID, p.Armazenamento)

	for {
		select {
		case err := <-ErrChan:
			p.ErrChan <- fmt.Errorf("processoID[%s]: %w", p.ProcessoID, err)
			// return TODO colocar o return aqui ou chamar o stop?

		case <-p.stopChan:
			// SIGINT ou SIGTERM o processo acima... gracefully shutdown
			close(stopCmdChan)
			return
		}
	}
}

// TODO vai ser a funcao no package Dahua
func dahua(outChan chan registro.Registro, ErrChan chan error, stopCmdChan chan struct{}, processoID string, armazenamento string) {
	var i int
	for {
		select {
		default:
			time.Sleep(time.Duration(time.Millisecond * 500))
			r := registro.Registro{
				RegistroID:    validate.GenerateID(), // TODO gerar dentro do dahua
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
			outChan <- r
			i++
		case <-stopCmdChan:
			return
		}
	}
}
