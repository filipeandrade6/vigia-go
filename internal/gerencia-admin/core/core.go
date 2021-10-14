package core

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/core/usuario"
	"github.com/filipeandrade6/vigia-go/internal/gerencia-admin/config"
	"github.com/filipeandrade6/vigia-go/internal/gerencia-admin/service"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
)

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

	// ----

	time.Sleep(time.Duration(time.Second * 5))
	gerenciaClient := service.NewClientGerencia(fmt.Sprintf("%s:%d", cfg.Service.Address, cfg.Service.Port))

	usuarioID := "ce93c4ba-aec9-42fa-ba7c-85e712e4ade8"

	u := usuario.UpdateUsuario{
		Funcao: []string{"USER", "MANAGER"},
		Senha:  &wrapperspb.StringValue{Value: "secret"},
	}

	fmt.Println(usuarioID)

	if err := gerenciaClient.UpdateUsuario(usuarioID, u); err != nil {
		log.Fatalw("updating usuario", "ERROR", err)
	}

	return nil
}
