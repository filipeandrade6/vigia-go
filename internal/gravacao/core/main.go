package core

import (
	"expvar"
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/filipeandrade6/vigia-go/internal/gravacao/server"
	"google.golang.org/grpc"

	// "github.com/ardanlabs/service/app/services/sales-api/handlers"
	// "github.com/ardanlabs/service/business/sys/database"
	// "github.com/ardanlabs/service/business/sys/metrics"
	// "github.com/filipeandrade6/vigia-go/business/sys/auth"
	config "github.com/filipeandrade6/vigia-go/internal/config"
	// "github.com/filipeandrade6/vigia-go/internal/keystore"

	// "go.opentelemetry.io/otel"
	// "go.opentelemetry.io/otel/attribute"
	// "go.opentelemetry.io/otel/exporters/zipkin"
	// "go.opentelemetry.io/otel/sdk/resource"
	// "go.opentelemetry.io/otel/sdk/trace"
	// semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
)

type Gravacao struct {
	server *grpc.Server
	// client *client.GerenciaClient
	// models models.Models
}

func (g *Gravacao) Stop() {
	fmt.Println("Finalizando aplicação....")
	g.server.GracefulStop() // TODO colocar context e finalizar forçado com 30 seg
	fmt.Println("Bye.")
}

// TODO colocar gRPC Health Server https://gist.github.com/akhenakh/38dbfea70dc36964e23acc19777f3869

/*
Need to figure out timeouts for http service.
*/

// build is the git versin of this program. It is set using build flags in the makefile.
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
	fmt.Println(cfg)

	// =========================================================================
	// App Starting

	expvar.NewString("build").Set(build)
	log.Infow("starting service", "version", build)
	defer log.Infow("shutdown complete")

	log.Infow("startup", "config", cfg)

	// =========================================================================
	// Initialize authentication support

	// log.Infow("startup", "status", "initializing authentication support")

	// ks, err := keystore.NewFS(os.DirFS(cfg.Auth.KeysFolder))
	// if err != nil {
	// 	return fmt.Errorf("reading keys: %w", err)
	// }

	// auth, err := auth.New(cfg.Auth.ActiveKID, ks)
	// if err != nil {
	// 	return fmt.Errorf("constructing auth: %w", err)
	// }

	// =========================================================================
	// Start Database

	// log.Infow("startup", "status", "initializing database support", "host", cfg.DB.Host)

	// db, err := database.Open(database.Config{
	// 	User:         cfg.DB.User,
	// 	Password:     cfg.DB.Password,
	// 	Host:         cfg.DB.Host,
	// 	Name:         cfg.DB.Name,
	// 	MaxIdleConns: cfg.DB.MaxIdleConns,
	// 	MaxOpenConns: cfg.DB.MaxOpenConns,
	// 	DisableTLS:   cfg.DB.DisableTLS,
	// })
	// if err != nil {
	// 	return fmt.Errorf("connecting to db: %w", err)
	// }
	// defer func() {
	// 	log.Infow("shutdown", "status", "stopping database support", "host", cfg.DB.Host)
	// 	db.Close()
	// }()

	// =========================================================================
	// Start Tracing Support

	// WARNING: The current Init settings are using defaults which may not be
	// compatible with your project. Please review the documentation for
	// opentelemetry.

	// log.Infow("startup", "status", "initializing OT/Zipkin tracing support")

	// exporter, err := zipkin.New(
	// 	cfg.Zipkin.ReporterURI,
	// 	// zipkin.WithLogger(zap.NewStdLog(log)),
	// )
	// if err != nil {
	// 	return fmt.Errorf("creating new exporter: %w", err)
	// }

	// traceProvider := trace.NewTracerProvider(
	// 	trace.WithSampler(trace.TraceIDRatioBased(cfg.Zipkin.Probability)),
	// 	trace.WithBatcher(exporter,
	// 		trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
	// 		trace.WithBatchTimeout(trace.DefaultBatchTimeout),
	// 		trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
	// 	),
	// 	trace.WithResource(
	// 		resource.NewWithAttributes(
	// 			semconv.SchemaURL,
	// 			semconv.ServiceNameKey.String(cfg.Zipkin.ServiceName),
	// 			attribute.String("exporter", "zipkin"),
	// 		),
	// 	),
	// )

	// // I can only get this working properly using the singleton :(
	// otel.SetTracerProvider(traceProvider)
	// defer traceProvider.Shutdown(context.Background())

	// =========================================================================
	// Start Debug Service

	// log.Infow("startup", "status", "debug router started", "host", cfg.Web.DebugHost)

	// // The Debug function returns a mux to listen and serve on for all the debug
	// // related endpoints. This include the standard library endpoints.

	// // Construct the mux for the debug calls.
	// debugMux := handlers.DebugMux(build, log, db)

	// // Start the service listening for debug requests.
	// // Not concerned with shutting this down with load shedding.
	// go func() {
	// 	if err := http.ListenAndServe(cfg.Web.DebugHost, debugMux); err != nil {
	// 		log.Errorw("shutdown", "status", "debug router closed", "host", cfg.Web.DebugHost, "ERROR", err)
	// 	}
	// }()

	// =========================================================================
	// Start Service

	log.Infow("startup", "status", "initializing API support")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	g := &Gravacao{
		server: server.NovoServidorGravacao(),
	}

	// dbCfg := g.client.ConfigBancoDeDados()

	// _, err := database.NewPool(dbCfg)
	// if err != nil {
	// 	return err
	// }

	serverErrors := make(chan error, 1)

	// TODO ver abaixo, tem exemplo toda execução em contexto
	// https://gist.github.com/akhenakh/38dbfea70dc36964e23acc19777f3869
	go func() {
		lis, err := net.Listen(cfg.Service.GravServerConn, fmt.Sprintf("%s:%s", cfg.Service.GravServerAddr, cfg.Service.GravServerPort))
		if err != nil {
			log.Errorw("startup", "status", "could not open socket", cfg.Service.GravServerConn, cfg.Service.GravServerAddr, cfg.Service.GravServerPort, "ERROR", err)
		}

		log.Infow("startup", "status", "gRPC server started") // TODO add address
		serverErrors <- g.server.Serve(lis)
	}()

	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Infow("shutdown", "status", "shutdown started", "signal", sig)
		defer log.Infow("shutdown", "status", "shutdown complete", "signal", sig)

		g.server.GracefulStop()

		// Give outstanding requests a deadline for completion.
		// ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		// defer cancel()

		// // Asking listener to shutdown and shed load.
		// if err := api.Shutdown(ctx); err != nil {
		// 	api.Close()
		// 	return fmt.Errorf("could not stop server gracefully: %w", err)
		// }
	}

	return nil
}
