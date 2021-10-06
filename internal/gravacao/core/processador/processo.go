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

type Processo struct {
	processo.Processo
	camera.Camera

	servidorgravacao string
	armazenamento    string

	// ProcessoID         string        // identificador do processo
	// ServidorGravacaoID string        // identificador do servidor de gravacao
	// Camera             camera.Camera // configurações da câmera
	// Processador        int           // processador utilizado
	// Adaptador          int           // adaptador utilizado
	// Armazenamento      string        // local onde será salvo as imagens

	Status bool // true: executando, false: parado

	regChan     chan registro.Registro // channel onde será enviado os registros TODO no Processador
	errChan     chan error             // channel onde será enviado erro
	stopChan    chan struct{}          // channel de sinalização para parar o processamento
	stoppedChan chan struct{}          // channel de sinalização que o processamento foi parado
}

func NewProcesso(
	prc processo.Processo,
	cam camera.Camera,
	servidorgravacao string,
	armazenamento string,
	regChan chan registro.Registro,
	errChan chan error,
) *Processo {
	return &Processo{
		Processo:         prc,
		Camera:           cam,
		servidorgravacao: servidorgravacao,
		armazenamento:    armazenamento,
		regChan:          make(chan registro.Registro),
		errChan:          make(chan error),
		Status:           false,
	}
}

// func NewProcesso(
// 	processoID string,
// 	servidorGravacaoID string,
// 	armazenamento string,
// 	processador int,
// 	adaptador int,
// 	camera camera.Camera,
// 	regChan chan registro.Registro,
// 	errChan chan error,
// ) *Processo {
// 	return &Processo{
// 		ProcessoID:         processoID,
// 		ServidorGravacaoID: servidorGravacaoID,
// 		Armazenamento:      armazenamento,
// 		Processador:        processador,
// 		Adaptador:          adaptador,
// 		Camera:             camera,
// 		regChan:            regChan,
// 		errChan:            errChan,
// 	}
// }

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

	errChan := make(chan error)
	stopCmdChan := make(chan struct{})

	go dahua(p.regChan, errChan, stopCmdChan, p.ProcessoID, p.armazenamento)

	for {
		select {
		case err := <-errChan:
			p.errChan <- fmt.Errorf("processoID[%s]: %w", p.ProcessoID, err)
			// return TODO colocar o return aqui ou chamar o stop?

		case <-p.stopChan:
			// SIGINT ou SIGTERM o processo acima... gracefully shutdown
			close(stopCmdChan)
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
