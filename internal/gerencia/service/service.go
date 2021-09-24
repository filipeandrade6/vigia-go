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

	// PERFORM PRE BUSINESS OPERATIONS

	cameraID, err := g.cameraStore.Create(ctx, cam, now)
	if err != nil {
		return "", fmt.Errorf("create: %w", err)
	}

	// PERFORM POST BUSINESS OPERATIONS

	return cameraID, nil
}

func (g *GerenciaService) ReadCameras(ctx context.Context, pageNumber int, rowsPerPage int) ([]camera.Camera, error) {

	// PERFORM PRE BUSINESS OPERATIONS

	cameras, err := g.cameraStore.Query(ctx, pageNumber, rowsPerPage)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	// PERFORM POST BUSINESS OPERATIONS

	return cameras, nil
}

func (g *GerenciaService) ReadCamera(ctx context.Context, cameraID string) (camera.Camera, error) {

	// PERFORM PRE BUSINESS OPERATIONS

	c, err := g.cameraStore.QueryByID(ctx, cameraID)
	if err != nil {
		return camera.Camera{}, fmt.Errorf("query: %w", err)
	}

	// PERFORM POST BUSINESS OPERATIONS

	return c, nil
}

func (g *GerenciaService) UpdateCamera(ctx context.Context, cam camera.Camera, now time.Time) error {

	// PERFORM PRE BUSINESS OPERATIONS:

	if err := g.cameraStore.Update(ctx, cam, now); err != nil {
		return fmt.Errorf("update: %w", err)
	}

	// PERFORM POST BUSINESS OPERATIONS

	return nil
}

func (g *GerenciaService) DeleteCamera(ctx context.Context, cameraID string) error {

	// PERFORM PRE BUSINESS OPERATIONS

	if err := g.cameraStore.Delete(ctx, cameraID); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	// PERFORM POST BUSINESS OPERATIONS

	return nil
}
