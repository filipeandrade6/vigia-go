package service

import (
	"github.com/filipeandrade6/vigia-go/internal/data/store/camera"
	"github.com/filipeandrade6/vigia-go/internal/data/store/processo"
	"github.com/filipeandrade6/vigia-go/internal/data/store/servidorgravacao"
	"go.uber.org/zap"
)

type GerenciaService struct {
	log                   *zap.SugaredLogger
	cameraStore           camera.Store
	processoStore         processo.Store
	servidorGravacaoStore servidorgravacao.Store
	// publisher
	// gravacaoClient *client.GravacaoClient
}

func NewGerenciaService(log *zap.SugaredLogger, cameraStore camera.Store, processoStore processo.Store, servidorGravacaoStore servidorgravacao.Store) *GerenciaService {
	return &GerenciaService{
		log:                   log,
		cameraStore:           cameraStore,
		processoStore:         processoStore,
		servidorGravacaoStore: servidorGravacaoStore,
		// gravacaoClient:        gravacaoClient,
	}
}
