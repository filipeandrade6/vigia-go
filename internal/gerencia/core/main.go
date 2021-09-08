// Core compreende as regras de negócio
package core

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "github.com/filipeandrade6/vigia-go/internal/database"
	"github.com/filipeandrade6/vigia-go/internal/gerencia/client"
	// "github.com/filipeandrade6/vigia-go/internal/gerencia/server"
	// "google.golang.org/grpc"
)

type Gerencia struct {
	// server *grpc.Server
	client *client.GravacaoClient // TODO verificar se campo privado não interfere em algo
}

// Stop para a aplicação depois de recebido um sinal de interrupção do sistema
// func (g *Gerencia) Stop() {
// 	fmt.Println("Finalizando aplicação....")
// 	g.server.GracefulStop() // TODO colocar context e finalizar forçado com 30 seg
// 	fmt.Println("Bye.")
// }

// Main é a funcao principal que inicia o server e client da API que intercomunica
// os servicos dos servidores de gerencia e gravacao
func Run() error {

	// logger,  := zap.NewProduction()
	// defer logger.Sync()

	g := &Gerencia{
		// server: server.NovoServidorGerencia(),
		client: client.NovoClientGravacao(),
	}

	// dbCfg := database.NewConfig()
	// _, err := database.NewPool(dbCfg) // TODO implementar
	// if err != nil {
	// 	return err
	// }

	fmt.Println("chegou aqui 1 gerenciaaaa")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("chegou aqui 2")

	time.Sleep(time.Duration(time.Second * 10))
	resp, err := g.client.IniciarProcessamento(ctx, nil)
	if err != nil {

		fmt.Println("deu erro", err)
	}
	fmt.Printf("chegou status: %s", resp.Status)

	fmt.Println("chegou aqui 3")
	<-c
	// g.Stop()

	fmt.Println("chegou aqui 4")
	return errors.New("finalizou funcao main")
}
