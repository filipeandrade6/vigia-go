package core

import (
	"expvar"
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/config"
	grpc_gravacao "github.com/filipeandrade6/vigia-go/internal/gravacao/grpc"

	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var build = "develop"

func Run(log *zap.SugaredLogger, cfg config.Configuration) error {
	// =========================================================================
	// CPU Quota

	if _, err := maxprocs.Set(); err != nil {
		log.Errorw("startup", zap.Error(err))
		os.Exit(1)
	}
	log.Infow("startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// =========================================================================
	// Show Configuration

	log.Infow("startup", "config", fmt.Sprintf("%+v", cfg)) // TODO esconder senhas

	// =========================================================================
	// App Starting

	expvar.NewString("build").Set(build)
	log.Infow("starting service", "version", build)
	defer log.Infow("shutdown complete")

	// =========================================================================
	// TODO Initialize Authentication Support

	// =========================================================================
	// Start gRPC Service

	log.Infow("startup", "status", "initializing API support")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	serverErrors := make(chan error, 1)

	svc := grpc_gravacao.NewGravacaoService(log, cfg.Gravacao.Armazenamento, cfg.Gravacao.Housekeeper)

	grpcServer := grpc.NewServer()
	pb.RegisterGravacaoServer(grpcServer, svc)

	go func() {
		lis, err := net.Listen(cfg.Gravacao.Conn, fmt.Sprintf(":%d", cfg.Gravacao.Port))
		if err != nil {
			log.Errorw("startup", "status", "could not open socket", cfg.Gravacao.Conn, cfg.Gravacao.Port, "ERROR", err)
		}

		log.Infow("startup", "status", "gRPC server started") // TODO adicionar host
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
