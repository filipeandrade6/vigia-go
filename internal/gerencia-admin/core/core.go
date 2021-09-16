package core

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"

	gerenciaGRPC "github.com/filipeandrade6/vigia-go/internal/gerencia/grpc"
	"github.com/filipeandrade6/vigia-go/internal/sys/config"
)

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

	// ----

	fmt.Println("chegou aqui antes de criar o client de genrecia")
	time.Sleep(time.Duration(time.Second * 10))
	gerenciaClient := gerenciaGRPC.NovoClientGerencia()
	fmt.Println("chegou aqui")
	if err := gerenciaClient.Migrate(); err != nil {
		log.Fatalf("calling migrate: %w", err)
	}

	return nil
}
