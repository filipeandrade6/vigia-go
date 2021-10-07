package processador

import (
	"errors"
	"fmt"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/core/camera"
	"github.com/filipeandrade6/vigia-go/internal/core/processo"
	"github.com/filipeandrade6/vigia-go/internal/core/registro"
	"github.com/filipeandrade6/vigia-go/internal/sys/validate"
)

var (
	ErrAlreadyStarted   = errors.New("processo already started")
	ErrAlreadyStopped   = errors.New("processo already stopped")
	ErrWrongCredentials = errors.New("user/password wrong")
	ErrCameraOffline    = errors.New("camera offline")
)

type Processo struct {
	processo.Processo
	camera.Camera

	servidorGravacaoID string
	armazenamento      string

	status bool

	regChan     chan registro.Registro
	errChan     chan error
	stopChan    chan struct{}
	stoppedChan chan struct{}
}

func NewProcesso(
	prc processo.Processo,
	cam camera.Camera,
	servidorGravacaoID string,
	armazenamento string,
	regChan chan registro.Registro,
	errChan chan error,
) *Processo {
	return &Processo{
		Processo:           prc,
		Camera:             cam,
		servidorGravacaoID: servidorGravacaoID,
		armazenamento:      armazenamento,
		regChan:            regChan,
		errChan:            make(chan error),
		status:             false,
	}
}

func (p *Processo) Start() {
	p.stopChan = make(chan struct{})
	p.stoppedChan = make(chan struct{})

	go p.processar()

	p.status = true
}

func (p *Processo) Stop() {
	close(p.stopChan)
	<-p.stoppedChan

	p.status = false
}

func (p *Processo) processar() {
	defer close(p.stoppedChan)

	errChan := make(chan error)
	stopCmdChan := make(chan struct{})

	go dahua(p.regChan, errChan, stopCmdChan, p.ProcessoID, p.armazenamento)

	for {
		select {
		case err := <-errChan:
			p.errChan <- fmt.Errorf("processoID[%s]: %w", p.ProcessoID, err)
			p.status = false
			return

		case <-p.stopChan:
			close(stopCmdChan) // TODO SIGINT ou SIGTERM o processo acima... gracefully shutdown
			return
		}
	}
}

// TODO vai ser a funcao no package Dahua
func dahua(outChan chan registro.Registro, errChan chan error, stopCmdChan chan struct{}, processoID string, armazenamento string) {
	var i int
	for {
		select {
		default:
			time.Sleep(time.Duration(time.Millisecond * 200))
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
