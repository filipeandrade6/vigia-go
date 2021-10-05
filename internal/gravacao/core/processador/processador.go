package processador

import (
	"fmt"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/core/registro"
)

type Gerenciador struct {
	processer   *Processer
	notificator *Notificator

	RegChan chan Registro
	ErrChan chan error
	NotChan chan string
}

func NewGereciador(svID string) *Gerenciador {
	prcChan := make(chan Processo)
	regChan := make(chan Registro)
	notChan := make(chan string)
	registroCore := registro.NewCore()

	return &Gerenciador{
		processer:   NewProcesser(svID, registroCore, prcChan, notChan),
		notificator: NewNotificator(svID, notChan),
	}
}

func (g *Gerenciador) Gerenciar() {
	for {
		select {
		case proc := <-d.Novo:
			d.NovoProcesso(proc)

		case reg := <-d.RegChan:
			d.registroCore.CreateRegistro(reg)

		case err := <-d.ErrChan:
			d.notification.Send(err)
		}
	}
}

// =================================================================================

type Notificator struct {
	servidorGravacaoID string
	notChan            chan string
}

func NewNotificator(svID string, notChan chan string) *Notificator {
	return &Notificator{
		servidorGravacaoID: svID,
		notChan:            notChan,
	}
}

func (n *Notificator) Send(err error) {
	fmt.Errorf("notificar: %w", err)
}

// =================================================================================

type Processer struct {
	ServidorGravacaoID string
	RegistroCore       *registro.Core
	PrcChan            chan Processo
	NotChan            chan string

	processos   map[string]Processo // TODO pointer e todo o resto?
	matchlist   map[string]bool
	errChan     chan error
	regChan     chan Registro
	notificator Notificator
}

// TODO como o adaptador vai funcionar...

func NewProcesser(servidorGravacaoID string, registroCore registro.Core, prcChan chan Processo, notChan chan string) *Processer {
	return &Processer{
		ServidorGravacaoID: servidorGravacaoID,
		RegistroCore:       &registroCore,
		PrcChan:            prcChan,
		NotChan:            notChan,

		processos: make(map[string]Processo),
		matchlist: make(map[string]bool),
		errChan:   make(chan error),
	}
}
 func (p *Proces)

func (p *Processer) NovoProcesso() string {}

func (p *Processer) Salvar(processo Processo) {
	// p.RegistroCore.Create()
}

// func (d *Processer) NovoProcesso(cameraID, cameraIP, usuario, senha, armazenamento string, canal, processador, adaptador int) string { // adicionar argumentos
// 	p := Processo{
// 		ProcessoID: validate.GenerateID(),
// 		ServidorGravacaoID: d.ServidorGravacaoID,
// 		CameraID: cameraID,
// 		Processador: processador,
// 		Adaptador: adaptador,
// 		Armazenamento: armazenamento,

// 		RegChan:   d.RegChan,
// 		ErrorChan: d.ErrChan,
// 		StopChan:  make(chan struct{}),
// 		StoppedChan: make(chan struct{}),
// 	}

// 	go processar(p)

// 	d.Processos[p.ProcessoID] = p

// 	return p.ProcessoID
// }

// =================================================================================

func processar(p Processo) { // adicionar
	defer close(stoppedChan)

	// stdOut := make(chan string)
	// stdErr := make(chan err/string)

	// ======== comando
	// processo_id
	// camera_ip
	// porta
	// canal
	// usuario
	// senha
	// armazenamento
	// timestamp

	for {
		select {
		case reg := <- stdOut: // TODO mandar direto o regChan filtra?
			// check se é mensagem de registro
			// etc
			regChan <- reg

		case err := <- stdErr:
			p.ErrorChan <- err
			return

		case <-p.StopChan:
			// SIGINT ou SIGTERM o processo acima... gracefully shutdown
			return
		}
	}
}

// =================================================================================

type Processo struct {
	ProcessoID         string // identificador do processo
	ServidorGravacaoID string // identificador do servidor de gravacao
	Camera             string // configurações da câmera
	Processador        int    // processador utilizado
	Adaptador          int    // adaptador utilizado
	Armazenamento      string // local onde será salvo as imagens

	Status bool // true: executando, false: parado

	RegChan     chan Registro // channel onde será enviado os registros TODO no Processer
	ErrorChan   chan error    // channel onde será enviado erro
	StopChan    chan struct{} // channel de sinalização para parar o processamento
	StoppedChan chan struct{} // channel de sinalização que o processamento foi parado
}

func (p *Processo) Stop() error {
	// TODO colocar context?
	close(p.StopChan)
	<-p.stoppedChan
	return nil
}

type Registro struct {
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
