// Core compreende as regras de negócio
package core

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/database"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/client"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/server"

	"google.golang.org/grpc"
)

type Gravacao struct {
	closed chan struct{}
	wg     sync.WaitGroup
	ticker *time.Ticker

	server *grpc.Server
	client *client.GerenciaClient
}

func (g *Gravacao) Run() {
	for {
		select {
		case <-g.closed:
			fmt.Println("close.....")
		case <-g.ticker.C:
			g.wg.Add(1)
			fmt.Println("add handle....")
			go handle(g)
		}
	}
}

func (g *Gravacao) Stop() {
	fmt.Println("recebido sinal de finalização......")
	close(g.closed)
	g.wg.Wait() // wait for all go goroutines to exit
	fmt.Println("tudo finalizado...")
}

func handle(gravacao *Gravacao) {
	defer gravacao.wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Print("#")
		time.Sleep(time.Second * 2)
	}
}

// Main e a funcao principal que inicia o server e client da API que intercomunica
// os servicos dos servidores de gerencia e gravacao
func Main() error {

	// logger,  := zap.NewProduction()
	// defer logger.Sync()

	g := &Gravacao{
		closed: make(chan struct{}),
		wg:     sync.WaitGroup{},
		ticker: time.NewTicker(time.Second * 2),
		server: server.StartServer("tcp", "localhost:12346"),
		client: client.GetGerenciaClient("localhost:12347"),
	}

	dbCfg := g.client.GetDatabase()

	_, err := database.NewPool(dbCfg)
	if err != nil {
		return err
	}

	c := make(chan os.Signal) // OS termination signal
	signal.Notify(c, os.Interrupt)

	go g.Run()

	// go alguma metodo que nunca finaliza
	select {
	case sig := <-c:
		fmt.Printf("Got %s signal. Finalizando...\n", sig)
		// metodo de parada graciosa
	}

	return errors.New("finalizou funcao main")
}
