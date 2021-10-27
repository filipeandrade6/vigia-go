package processador

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ardanlabs/service/business/sys/validate"
	"github.com/filipeandrade6/vigia-go/internal/core/registro"
	"github.com/filipeandrade6/vigia-go/internal/sys/operrors"
)

type Camera interface {
	// New(processoID, enderecoIP string, porta, canal int, usuario, senha string)
	GetID() string
	Start(armazenamento string, regChan chan registro.Registro, errChan chan *operrors.OpError)
	Stop()
}

// type Processo struct {
// 	ProcessoID    string
// 	EnderecoIP    string
// 	Porta         int
// 	Canal         int
// 	Usuario       string
// 	Senha         string
// 	Processador   int
// 	Armazenamento string
// 	regChan       chan registro.Registro
// 	errChan       chan *operrors.OpError
// 	stopChan      chan struct{}
// 	stoppedChan   chan struct{}
// }

// func NewProcesso(
// 	processoID string,
// 	enderecoIP string,
// 	porta int,
// 	canal int,
// 	usuario string,
// 	senha string,
// 	processador int,
// 	armazenamento string,
// 	regChan chan registro.Registro,
// 	errChan chan *operrors.OpError,
// ) *Processo {
// 	return &Processo{
// 		ProcessoID:    processoID,
// 		EnderecoIP:    enderecoIP,
// 		Porta:         porta,
// 		Canal:         canal,
// 		Usuario:       usuario,
// 		Senha:         senha,
// 		Processador:   processador,
// 		Armazenamento: armazenamento,
// 		regChan:       regChan,
// 		errChan:       errChan,
// 	}
// }

// func (p *Processo) Start() {
// 	p.stopChan = make(chan struct{})
// 	p.stoppedChan = make(chan struct{})

// 	if p.Processador == 0 {
// 		go processoTeste(
// 			p.ProcessoID,
// 			p.EnderecoIP,
// 			p.Porta,
// 			p.Canal,
// 			p.Usuario,
// 			p.Senha,
// 			p.Armazenamento,
// 			p.regChan,
// 			p.errChan,
// 			p.stopChan,
// 			p.stoppedChan,
// 		)
// 	} else {
// 		go traffic.Start(
// 			p.ProcessoID,
// 			p.Armazenamento,
// 			p.EnderecoIP,
// 			int32(p.Porta),
// 			int32(p.Canal),
// 			p.Usuario,
// 			p.Senha,
// 			p.regChan,
// 			p.errChan,
// 			p.stopChan,
// 			p.stoppedChan,
// 		)
// 	}
// }

// func (p *Processo) Stop() {
// 	close(p.stopChan)
// 	<-p.stoppedChan
// }

type CameraTeste struct {
	ProcessoID    string
	EnderecoIP    string
	Porta         int
	Canal         int
	Usuario       string
	Senha         string
	Processador   int
	Armazenamento string
	stopChan      chan struct{}
	stoppedChan   chan struct{}
}

func NewCameraTeste(
	processoID string,
	enderecoIP string,
	porta int,
	canal int,
	usuario string,
	senha string,
) *CameraTeste {
	return &CameraTeste{
		ProcessoID: processoID,
		EnderecoIP: enderecoIP,
		Porta:      porta,
		Canal:      canal,
		Usuario:    usuario,
		Senha:      senha,
	}
}

func (c *CameraTeste) Start(armazenamento string, regChan chan registro.Registro, errChan chan *operrors.OpError) {
	c.stopChan = make(chan struct{})
	c.stoppedChan = make(chan struct{})

	go c.start(
		armazenamento,
		regChan,
		errChan,
	)
}

func (c *CameraTeste) Stop() {
	close(c.stopChan)
	<-c.stoppedChan
}

func (c *CameraTeste) GetID() string {
	return c.ProcessoID
}

// func processoTeste(
// 	processoID string,
// 	enderecoIP string,
// 	porta int,
// 	canal int,
// 	usuario string,
// 	senha string,
// 	armazenamento string,
// 	regChan chan registro.Registro,
// 	errChan chan *operrors.OpError,
// 	stopChan chan struct{},
// 	stoppedChan chan struct{},
// ) {

func (c *CameraTeste) start(armazenamento string, regChan chan registro.Registro, errChan chan *operrors.OpError) {
	defer close(c.stoppedChan)

	var i int
	for {
		select {
		case <-c.stopChan:
			fmt.Println("cancel called")
			return

		default:
			fmt.Print(i, "..")
			time.Sleep(time.Duration(time.Millisecond * 500))
			r := registro.Registro{
				RegistroID:    validate.GenerateID(),
				ProcessoID:    c.ProcessoID,
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

			err := os.WriteFile(filepath.Join(armazenamento, fmt.Sprintf("%d-%s.txt", i, r.RegistroID)), []byte("hello\ngo\n"), 0644)
			if err != nil {
				errChan <- &operrors.OpError{ProcessoID: c.ProcessoID, Err: err}
			}
			i++
		}
	}
}
