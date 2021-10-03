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
	"github.com/filipeandrade6/vigia-go/internal/core/camera"
	"github.com/filipeandrade6/vigia-go/internal/core/servidorgravacao"
	"github.com/filipeandrade6/vigia-go/internal/core/usuario"
	"github.com/filipeandrade6/vigia-go/internal/gerencia/config"
	"github.com/filipeandrade6/vigia-go/internal/gerencia/service"
	"github.com/filipeandrade6/vigia-go/internal/sys/auth"
	"github.com/filipeandrade6/vigia-go/internal/sys/database"
	"github.com/filipeandrade6/vigia-go/internal/sys/keystore"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
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
	// Load Configuration

	log.Infow("startup", "config", fmt.Sprintf("%+v", cfg)) // TODO esconder senhas

	// =========================================================================
	// App Starting

	expvar.NewString("build").Set(build)
	log.Infow("starting service", "version", build)
	defer log.Infow("shutdown complete")

	// =========================================================================
	// Authentication Support

	log.Infow("startup", "status", "initializing authentication support")

	ks, err := keystore.NewFS(os.DirFS(cfg.Auth.Directory))
	if err != nil {
		return fmt.Errorf("reading keys: %w", err)
	}

	auth, err := auth.New(cfg.Auth.ActiveKID, ks)
	if err != nil {
		return fmt.Errorf("constructing auth: %w", err)
	}

	// =========================================================================
	// Database Support

	log.Infow("startup", "status", "initializing database support", "host", cfg.Database.Host)

	db, err := database.Open(database.Config{
		User:         cfg.Database.User,
		Password:     cfg.Database.Password,
		Host:         cfg.Database.Host,
		Name:         cfg.Database.Name,
		MaxIDLEConns: cfg.Database.MaxIDLEConns,
		MaxOpenConns: cfg.Database.MaxOpenConns,
		SSLMode:      cfg.Database.SSLMode,
	})
	if err != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}
	defer func() {
		log.Infow("shutdown", "status", "stopping database support", "host", cfg.Database.Host)
		db.Close()
	}()

	// =========================================================================
	// TODO Start Tracing Support

	// =========================================================================
	// Start Debug Service

	// log.Infow("startup", "status", "debug router started", "host", viper.GetString("VIGIA_GER_DEBUGHOST"))

	// The Debug function returns a mux to listen and serve on for all the debug
	// related endpoints. This include the standart library endpoints.

	// Construct the mux for the debug calls.
	// debugMux := handlers.DebugMux(build, log, db)

	// Start the service listening for debug requests.
	// Not concerned with shutting this down with load shedding.
	// go func() {
	// 	if err := http.ListenAndServe(viper.GetString("VIGIA_GER_DEBUGHOST"), debugMux); err != nil {
	// 		log.Errorw("shutdown", "status", "debug router closed", "host", viper.GetString("VIGIA_GER_DEBUGHOST"), "ERROR", err)
	// 	}
	// }()

	// =========================================================================
	// Start Service

	log.Infow("startup", "status", "initializing gerencia")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	serverErrors := make(chan error, 1)

	cameraCore := camera.NewCore(log, db)
	usuarioCore := usuario.NewCore(log, db)
	servidorgravacaoCore := servidorgravacao.NewCore(log, db)

	svc := service.NewFrontendService(
		log,
		auth,
		database.Config{
			User:         cfg.Database.User,
			Password:     cfg.Database.Password,
			Host:         cfg.Database.Host,
			Name:         cfg.Database.Name,
			MaxIDLEConns: cfg.Database.MaxIDLEConns,
			MaxOpenConns: cfg.Database.MaxOpenConns,
			SSLMode:      cfg.Database.Host,
		},
		cameraCore,
		usuarioCore,
		servidorgravacaoCore,
	)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(nil)),
	)

	pb.RegisterFrontendServer(grpcServer, svc)

	go func() {
		lis, err := net.Listen(cfg.Service.Conn, fmt.Sprintf(":%d", cfg.Service.Port))
		if err != nil {
			log.Errorw("startup", "status", "could not open socket", cfg.Service.Conn, cfg.Service.Port, "ERROR", err)
		}

		log.Infow("startup", "status", "gRPC server started", cfg.Service.Address)
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
