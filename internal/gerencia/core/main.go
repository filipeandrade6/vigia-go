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
	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/config"
	"github.com/filipeandrade6/vigia-go/internal/data/store/camera"
	"github.com/filipeandrade6/vigia-go/internal/data/store/processo"
	"github.com/filipeandrade6/vigia-go/internal/data/store/servidorgravacao"

	// "github.com/filipeandrade6/vigia-go/internal/gerencia/client"
	gerenciaGRPC "github.com/filipeandrade6/vigia-go/internal/gerencia/grpc"
	"github.com/filipeandrade6/vigia-go/internal/gerencia/service"
	"github.com/filipeandrade6/vigia-go/internal/sys/database"
	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// TODO colocar metricas

type Gerencia struct {
	server *grpc.Server
	// client *client.GravacaoClient
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

	err := config.Load()
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	// log.Infow("startup", "config", cfg) // TODO criar um prettyprint para o cfg no log

	// =========================================================================
	// App Starting

	expvar.NewString("build").Set(build)
	log.Infow("starting service", "version", build)
	defer log.Infow("shutdown complete")

	// log.Infow("startup", "config", cfg)

	// =========================================================================
	// Start Database

	log.Infow("startup", "status", "initializing database support", "host", viper.GetString("DB_HOST"))

	db, err := database.Open(database.Config{
		Host:         viper.GetString("DB_HOST"),
		User:         viper.GetString("DB_USER"),
		Password:     viper.GetString("DB_PASSWORD"),
		Name:         viper.GetString("DB_NAME"),
		MaxIdleConns: viper.GetInt("DB_MAXIDLECONNS"),
		MaxOpenConns: viper.GetInt("DB_MAXOPENCONNS"),
		DisableTLS:   viper.GetBool("DB_DISABLETLS"),
	}) // TODO arrumar isso aqui - OPen não funcionando
	if err != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}
	defer func() {
		log.Infow("shutdown", "status", "stopping database support", "host", viper.GetString("DB_HOST"))
		db.Close()
	}()

	// =========================================================================
	// Start Tracing Support
	// Start Debug Service

	// Copiar do gravacao

	// =========================================================================
	// Start Service

	log.Infow("startup", "status", "initializing gerencia")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	serverErrors := make(chan error, 1)

	cameraStore := camera.NewStore(log, db)
	processoStore := processo.NewStore(log, db)
	servidorGravacaoStore := servidorgravacao.NewStore(log, db)
	// gravacaoClient := client.NovoClientGravacao() // TODO: passar log?
	svc := service.NewGerenciaService(log, cameraStore, processoStore, servidorGravacaoStore)

	grpcServer := grpc.NewServer()
	gerenciaGRPCService := gerenciaGRPC.NewGerenciaService(log, svc)

	pb.RegisterGerenciaServer(grpcServer, gerenciaGRPCService)

	// TODO ver abaixo, tem exemplo toda execução em contexto
	// https://gist.github.com/akhenakh/38dbfea70dc36964e23acc19777f3869
	go func() {
		lis, err := net.Listen(viper.GetString("GER_SERVER_CONN"), fmt.Sprintf(":%s", viper.GetString("GER_SERVER_PORT")))
		if err != nil {
			log.Errorw("startup", "status", "could not open socket", viper.GetString("GER_SERVER_CONN"), viper.GetString("GER_SERVER_ADDR"), viper.GetString("GER_SERVER_PORT"), "ERROR", err)
		}

		log.Infow("startup", "status", "gRPC server started") // TODO add address
		serverErrors <- grpcServer.Serve(lis)
	}()

	// =========================================================================
	// Shutdown

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Infow("shutdown", "status", "shutdown started", "signal", sig)
		defer log.Infow("shutdown", "status", "shutdown complete", "signal", sig)

		grpcServer.GracefulStop()
	}

	return nil
}
