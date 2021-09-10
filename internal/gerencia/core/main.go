// Core compreende as regras de negócio
package core

import (
	"expvar"
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	// "github.com/filipeandrade6/vigia-go/internal/database"
	"github.com/filipeandrade6/vigia-go/internal/config"
	"github.com/filipeandrade6/vigia-go/internal/gerencia/client"
	"github.com/filipeandrade6/vigia-go/internal/gerencia/server"
	"github.com/filipeandrade6/vigia-go/internal/sys/database"
	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// TODO colocar metricas

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
	// =========================================================================
	// CPU Quota

	if _, err := maxprocs.Set(); err != nil {
		log.Errorw("startup", zap.Error(err))
		os.Exit(1)
	}
	log.Infow("startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// =========================================================================
	// Configuration

	cfg, err := config.Load(build)
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	log.Infow("startup", "config", cfg) // TODO criar um prettyprint para o cfg no log

	// =========================================================================
	// App Starting

	expvar.NewString("build").Set(build)
	log.Infow("starting service", "version", build)
	defer log.Infow("shutdown complete")

	log.Infow("startup", "config", cfg)

	// =========================================================================
	// Start Database

	log.Infow("startup", "status", "initializing database support", "host", viper.GetString()  cfg.DB.Host)

	db, err := database.Open(database.Config{
		Host:         cfg.DB.Host,
		User:         cfg.DB.User,
		Password:     cfg.DB.Password,
		Name:         cfg.DB.Name,
		MaxIdleConns: cfg.DB.MaxIdleConns,
		MaxOpenConns: cfg.DB.MaxOpenConns,
		DisableTLS:   cfg.DB.DisableTLS,
	})
	if err != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}
	defer func() {
		log.Infow("shutdown", "status", "stopping database support", "host", cfg.DB.Host)
		db.Close()
	}()

	// =========================================================================
	// Start Service

	log.Infow("startup", "status", "initializing gerencia")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// TODO colocar client e database dentro do server?
	g := &Gerencia{
		server: server.NovoServidorGerencia(),
		client: client.NovoClientGravacao(),
		database: db,
	}

	g.server.Regis

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
}
