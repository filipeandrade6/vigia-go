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

	"github.com/filipeandrade6/vigia-go/internal/data/store/camera"
	"github.com/filipeandrade6/vigia-go/internal/gerencia-admin/service"
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
	gerenciaClient := service.NewClientGerencia()
	fmt.Println("chegou aqui")

	if err := gerenciaClient.Migrate(); err != nil {
		if errors.As(err, &migrate.ErrNoChange) {
			fmt.Println("Sem alterações")
		} else {
			log.Fatalf("calling migrate: %s", err)
		}
	}

	// TODO na migracao

	c := camera.Camera{
		Descricao:      "Camerasss 1",
		EnderecoIP:     "10.0.0.11",
		Porta:          12,
		Canal:          1,
		Usuario:        "admin",
		Senha:          "admin",
		Geolocalizacao: "-12.3242, -45.1234",
	}

	cam1, err := gerenciaClient.CreateCamera(c)
	if err != nil {
		log.Fatal(err)
	}

	c.Descricao = "Camerasss 2"
	c.EnderecoIP = "10.0.0.22"

	if _, err := gerenciaClient.CreateCamera(c); err != nil {
		log.Fatal(err)
	}

	c.Descricao = "Camerasss 3"
	c.EnderecoIP = "10.0.0.33"

	cam3, err := gerenciaClient.CreateCamera(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("criado 3 câmeras... segue Camerasss 2 abaixo")

	cRes, err := gerenciaClient.ReadCameras("asss 2", 1, 100)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cRes)

	c.CameraID = cam1
	c.Descricao = "Camera Updatadassss"
	c.EnderecoIP = "234.234.234.234"

	if err = gerenciaClient.UpdateCamera(c); err != nil {
		log.Fatal(err)
	}

	if err := gerenciaClient.DeleteCamera(cam3); err != nil {
		log.Fatal(err)
	}

	fmt.Println("câmera 3 deletada e câmera 1 atualizada... segue abaixo")

	cRes, err = gerenciaClient.ReadCameras("", 1, 100)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cRes)

	return nil
}
