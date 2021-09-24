package service

import (
	"context"
	"fmt"
	"time"

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

func (g *GerenciaService) CreateCamera(ctx context.Context, cam camera.Camera, now time.Time) (string, error) {

	c, err := g.cameraStore.Create(ctx, cam, now)
	if err != nil {
		return "", fmt.Errorf("create: %w", err)
	}

	return c.CameraID, nil
}

func (g *GerenciaService) ReadCamera(ctx context.Context, cameraID string) (camera.Camera, error) {

	c, err := g.cameraStore.QueryByID(ctx, cameraID)
	if err != nil {
		return camera.Camera{}, fmt.Errorf("query: %w", err)
	}

	return c, nil
}

func (g *GerenciaService) UpdateCamera(ctx context.Context, cam camera.Camera, now time.Time) {

}

func (g *GerenciaService) DeleteCamera(ctx context.Context, cam camera.Camera) {

}
