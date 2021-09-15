package core

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"

	"github.com/filipeandrade6/vigia-go/internal/gerencia/grpc/client"
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
	// Configuration

	viper.AutomaticEnv()

	// ----

	fmt.Println("chegou aqui antes de criar o client de genrecia")
	time.Sleep(time.Duration(time.Second * 10))
	gerenciaClient := client.NovoClientGerencia()
	fmt.Println("chegou aqui")
	if err := gerenciaClient.Migrate(); err != nil {
		log.Fatalf("calling migrate: %w", err)
	}

	return nil
}
