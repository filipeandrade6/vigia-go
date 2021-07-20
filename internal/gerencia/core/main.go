// Core compreende as regras de negócio
package core

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/filipeandrade6/vigia-go/internal/database"
	"github.com/filipeandrade6/vigia-go/internal/gerencia/client"
	"github.com/filipeandrade6/vigia-go/internal/gerencia/server"

	"google.golang.org/grpc"
)

type Gerencia struct {
	server *grpc.Server
	client *client.GravacaoClient // TODO verificar se campo privado não interfere em algo
}

func (g *Gerencia) Stop() {
	fmt.Println("Finalizando aplicação....")
	// g.server.GracefulStop() // TODO colocar context e finalizar forçado com 30 seg
	fmt.Println("Bye.")
}

// Main e a funcao principal que inicia o server e client da API que intercomunica
// os servicos dos servidores de gerencia e gravacao
func Main() error {

	// logger,  := zap.NewProduction()
	// defer logger.Sync()

	// TODO ser semantico nos nomes, definindo bem qual é client/server de que serviço
	g := &Gerencia{
		server: server.NovoServidorGerencia("tcp", "localhost:12347"),
		client: client.NovoClientGerencia("localhost:12346"),
	}

	dbCfg := g.client.GetDatabase()

	_, err := database.NewPool(dbCfg) // TODO arruma aqui
	if err != nil {
		return err
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c
	g.Stop()

	return errors.New("finalizou funcao main")
}
