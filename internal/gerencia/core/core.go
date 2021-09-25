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

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/data/store/camera"
	"github.com/filipeandrade6/vigia-go/internal/gerencia/service"
	"github.com/filipeandrade6/vigia-go/internal/sys/auth"
	"github.com/filipeandrade6/vigia-go/internal/sys/config"
	"github.com/filipeandrade6/vigia-go/internal/sys/database"
	"github.com/filipeandrade6/vigia-go/internal/sys/keystore"

	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

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
	// Load Configuration

	viper.AutomaticEnv()
	log.Infow("startup", "config", config.PrettyPrintConfig())

	// =========================================================================
	// App Starting

	expvar.NewString("build").Set(build)
	log.Infow("starting service", "version", build)
	defer log.Infow("shutdown complete")

	// =========================================================================
	// Authentication Support

	log.Infow("startup", "status", "initializing authentication support")

	ks, err := keystore.NewFS(os.DirFS(viper.GetString("VIGIA_AUTH_DIR")))
	if err != nil {
		return fmt.Errorf("reading keys: %w", err)
	}

	auth, err := auth.New(viper.GetString("VIGIA_AUTH_ACTIVEKID"), ks)
	if err != nil {
		return fmt.Errorf("constructing auth: %w", err)
	}

	// =========================================================================
	// Database Support

	log.Infow("startup", "status", "initializing database support", "host", viper.GetString("VIGIA_DB_HOST"))

	db, err := database.Open(database.Config{
		Host:         viper.GetString("VIGIA_DB_HOST"),
		User:         viper.GetString("VIGIA_DB_USER"),
		Password:     viper.GetString("VIGIA_DB_PASSWORD"),
		Name:         viper.GetString("VIGIA_DB_NAME"),
		MaxIdleConns: viper.GetInt("VIGIA_DB_MAXIDLECONNS"),
		MaxOpenConns: viper.GetInt("VIGIA_DB_MAXOPENCONNS"),
		SSLMode:      viper.GetString("VIGIA_DB_SSLMODE"),
	})
	if err != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}
	defer func() {
		log.Infow("shutdown", "status", "stopping database support", "host", viper.GetString("VIGIA_DB_HOST"))
		db.Close()
	}()

	// =========================================================================
	// TODO Start Tracing Support

	// =========================================================================
	// TODO Start Debug Service

	// =========================================================================
	// Start Service

	log.Infow("startup", "status", "initializing gerencia")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	serverErrors := make(chan error, 1)

	cameraStore := camera.NewStore(log, db)
	svc := service.NewGerenciaService(log, auth, cameraStore)

	grpcServer := grpc.NewServer()
	pb.RegisterGerenciaServer(grpcServer, svc)

	go func() {
		lis, err := net.Listen(viper.GetString("VIGIA_GER_SERVER_CONN"), fmt.Sprintf(":%s", viper.GetString("VIGIA_GER_SERVER_PORT")))
		if err != nil {
			log.Errorw("startup", "status", "could not open socket", viper.GetString("VIGIA_GER_SERVER_CONN"), viper.GetString("VIGIA_GER_SERVER_ADDR"), viper.GetString("VIGIA_GER_SERVER_PORT"), "ERROR", err)
		}

		log.Infow("startup", "status", "gRPC server started", viper.GetString("VIGIA_GER_SERVER_HOST"))
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
