package core

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"

	gerenciaGRPC "github.com/filipeandrade6/vigia-go/internal/gerencia/grpc"
	"github.com/filipeandrade6/vigia-go/internal/sys/config"
	"github.com/filipeandrade6/vigia-go/internal/data/store/camera"
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
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("Sem alterações")
		} else {
			log.Fatalf("calling migrate: %w", err)
		}
	}

	c := camera.Camera{
		Descricao:      "Camera 1",
		EnderecoIP:     "10.0.0.1",
		Porta:          12,
		Canal:          1,
		Usuario:        "admin",
		Senha:          "admin",
		Geolocalizacao: "-12.3242, -45.1234",
	}

	if err := gerenciaClient.CreateCamera(c); err != nil {
		log.Fatal(err)
	}

	c.Descricao = "Camera 2"
	c.EnderecoIP = "10.0.0.2"

	if err := gerenciaClient.CreateCamera(c); err != nil {
		log.Fatal(err)
	}

	c.Descricao = "Camera 3"
	c.EnderecoIP = "10.0.0.3"

	if err := gerenciaClient.CreateCamera(c); err != nil {
		log.Fatal(err)
	}

	if err := gerenciaClient.UpdateCamera




	return nil
}
