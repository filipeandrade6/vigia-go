// Core compreende as regras de negócio
package core

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/filipeandrade6/vigia-go/internal/database"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/client"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/server"

	"google.golang.org/grpc"
)

type Gravacao struct {
	server *grpc.Server
	client *client.GerenciaClient // TODO verificar se campo privado não interfere em algo
}

func (g *Gravacao) Stop() {
	fmt.Println("Finalizando aplicação....")
	g.server.GracefulStop() // TODO colocar context e finalizar forçado com 30 seg
	fmt.Println("Bye.")
}

// Main e a funcao principal que inicia o server e client da API que intercomunica
// os servicos dos servidores de gerencia e gravacao
func Main() error {

	// logger,  := zap.NewProduction()
	// defer logger.Sync()

	g := &Gravacao{
		server: server.NovoServidorGravacao(),
		client: client.NovoClientGerencia(),
	}

	dbCfg := g.client.GetDatabase()

	_, err := database.NewPool(dbCfg) // TODO arrumar aqui
	if err != nil {
		return err
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c
	g.Stop()

	return errors.New("finalizou funcao main")
}
