// Core compreende as regras de negócio
package core

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	// "github.com/filipeandrade6/vigia-go/internal/database"
	"github.com/filipeandrade6/vigia-go/internal/config"
	"github.com/filipeandrade6/vigia-go/internal/gerencia/client"
	"github.com/filipeandrade6/vigia-go/internal/gerencia/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Gerencia struct {
	server *grpc.Server
	client *client.GravacaoClient
}

func (g *Gerencia) Stop() {
	fmt.Println("Finalizando aplicação...")
	g.server.GracefulStop() // TODO colocar context e finalizar forçado com 30seg ou menos
	fmt.Println("Bye.")
}

// TODO colocar gRPC Health Server https://gist.github.com/akhenakh/38dbfea70dc36964e23acc19777f3869

var build = "develop"

func Run(log *zap.SugaredLogger) error {

	cfg, err := config.Load(build)
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	log.Infow("startup", "config", cfg) // TODO criar um prettyprint para o cfg no log

	// =========================================================================
	// Start Service

	log.Infow("startup", "status", "initializing gerencia")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	g := &Gerencia{
		server: server.NovoServidorGerencia(),
		client: client.NovoClientGravacao(),
	}

	serverErrors := make(chan error, 1)

	// TODO ver abaixo, tem exemplo toda execução em contexto
	// https://gist.github.com/akhenakh/38dbfea70dc36964e23acc19777f3869
	go func() {
		lis, err := net.Listen(cfg.Gerencia.ServerConn, fmt.Sprintf("%s:%s", cfg.Gerencia.ServerAddr, cfg.Gerencia.ServerPort))
		if err != nil {
			log.Errorw("startup", "status", "could not open socket", cfg.Gerencia.ServerConn, cfg.Gerencia.ServerAddr, cfg.Gerencia.ServerPort, "ERROR", err)
		}

		log.Infow("startup", "status", "gRPC server started") // TODO add address
		serverErrors <- g.server.Serve(lis)
	}()

	// =========================================================================
	// Shutdown

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Infow("shutdown", "status", "shutdown started", "signal", sig)
		defer log.Infow("shutdown", "status", "shutdown complete", "signal", sig)

		g.server.GracefulStop()
	}

	return nil
	// fmt.Println("chegou aqui 1 gerenciaaaa")

	// c := make(chan os.Signal, 1)
	// signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// fmt.Println("chegou aqui 2")

	// time.Sleep(time.Duration(time.Second * 10))
	// resp, err := g.client.IniciarProcessamento(ctx, nil)
	// if err != nil {

	// 	fmt.Println("deu erro", err)
	// }
	// fmt.Printf("chegou status: %s", resp.Status)

	// fmt.Println("chegou aqui 3")
	// <-c
	// // g.Stop()

	// fmt.Println("chegou aqui 4")
	// return errors.New("finalizou funcao main")
}
