package core

import (
	"expvar"
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	pb "github.com/filipeandrade6/vigia-go/internal/api"
	gerenciaGRPC "github.com/filipeandrade6/vigia-go/internal/gerencia/grpc"
	gravacaoGRPC "github.com/filipeandrade6/vigia-go/internal/gravacao/grpc"
	gravacaoService "github.com/filipeandrade6/vigia-go/internal/gravacao/service"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	// "github.com/ardanlabs/service/app/services/sales-api/handlers"
	"github.com/filipeandrade6/vigia-go/internal/sys/config"

	// "github.com/ardanlabs/service/business/sys/metrics"
	// "github.com/filipeandrade6/vigia-go/business/sys/auth"

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
	// Load Configuration

	viper.AutomaticEnv()
	log.Infow("startup", "config", config.PrettyPrintConfig())

	// =========================================================================
	// App Starting

	expvar.NewString("build").Set(build)
	log.Infow("starting service", "version", build)
	defer log.Infow("shutdown complete")

	// =========================================================================
	// TODO Initialize Authentication Support

	// =========================================================================
	// Start Database
	// TODO database.Open não funciona

	// log.Infow("startup", "status", "initializing database support", "host", viper.GetString("DB_HOST"))

	// db, err := database.Open(database.Config{
	// 	Host:         viper.GetString("DB_HOST"),
	// 	User:         viper.GetString("DB_USER"),
	// 	Password:     viper.GetString("DB_PASSWORD"),
	// 	Name:         viper.GetString("DB_NAME"),
	// 	MaxIdleConns: viper.GetInt("DB_MAXIDLECONNS"),
	// 	MaxOpenConns: viper.GetInt("DB_MAXOPENCONNS"),
	// 	DisableTLS:   viper.GetBool("DB_DISABLETLS"),
	// })
	// if err != nil {
	// 	return fmt.Errorf("connecting to db: %w", err)
	// }
	// defer func() {
	// 	log.Infow("shutdown", "status", "stopping database support", "host", viper.GetString("DB_HOST"))
	// 	db.Close()
	// }()

	// =========================================================================
	// TODO Start Tracing Support

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
	// TODO Start Debug Service

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

	serverErrors := make(chan error, 1)

	gerenciaClient := gerenciaGRPC.NovoClientGerencia() // TODO: passar log?
	svc := gravacaoService.NewGravacaoService(log, gerenciaClient)

	grpcServer := grpc.NewServer()
	gravacaoGRPCService := gravacaoGRPC.NewGravacaoService(log, svc)

	pb.RegisterGravacaoServer(grpcServer, gravacaoGRPCService)

	go func() {
		lis, err := net.Listen(viper.GetString("VIGIA_GRA_SERVER_CONN"), fmt.Sprintf(":%s", viper.GetString("VIGIA_GRA_SERVER_PORT")))
		if err != nil {
			log.Errorw("startup", "status", "could not open socket", viper.GetString("VIGIA_GRA_SERVER_CONN"), viper.GetString("VIGIA_GRA_SERVER_PORT"), "ERROR", err)
		}

		log.Infow("startup", "status", "gRPC server started", viper.GetString("VIGIA_GRA_SERVER_HOST"))
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
