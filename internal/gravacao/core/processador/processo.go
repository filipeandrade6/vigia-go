package processador

import (
	"errors"
	"fmt"
	"time"

	"github.com/ardanlabs/service/business/sys/validate"
	"github.com/filipeandrade6/vigia-go/internal/core/camera"
	"github.com/filipeandrade6/vigia-go/internal/core/registro"
)

type Registro struct {
	RegistroID         string
	ProcessoID         string
	CameraID           string
	ServidorGravacaoID string
	Placa              string
	CorVeiculo         string
	TipoVeiculo        string
	MarcaVeiculo       string
	Armazenamento      string
	Horario            time.Time
}

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
	servidorGravacaoID string,
	armazenamento string,
	processador int,
	adaptador int,
	camera camera.Camera,
	regChan chan registro.Registro,
	ErrChan chan error) *Processo {
	return &Processo{
		ProcessoID:         validate.GenerateID(),
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
		return errors.New("processo ja em execucao")
	}

	p.stopChan = make(chan struct{})
	p.stoppedChan = make(chan struct{})

	go p.processar()
	p.Status = true

	return nil
}

func (p *Processo) Stop() error {
	if !p.Status {
		return errors.New("processo ja em pausa")
	}

	// TODO colocar context?
	close(p.stopChan)
	<-p.stoppedChan

	p.Status = false

	return nil
}

func (p *Processo) processar() {
	defer close(p.stoppedChan)

	outChan := make(chan string)
	ErrChan := make(chan error)

	stopCmdChan := make(chan struct{})

	go dahua(outChan, ErrChan, stopCmdChan)

	for {
		select {
		case r := <-outChan:
			// TODO check se é mensagem de registro adaptador?
			reg := registro.Registro{
				ProcessoID:   p.ProcessoID,
				Placa:        r,
				TipoVeiculo:  "sedan",
				CorVeiculo:   "prata",
				MarcaVeiculo: "honda",
				Confianca:    0.50,
				CriadoEm:     time.Now(),
			}

			p.regChan <- reg

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
func dahua(outChan chan string, ErrChan chan error, stopCmdChan chan struct{}) {
	outChan <- "TESTE0001"
	ErrChan <- errors.New("erro na dahua")
	<-stopCmdChan
}
