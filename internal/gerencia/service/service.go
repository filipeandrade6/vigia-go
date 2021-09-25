package service

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/data/migration"
	"github.com/filipeandrade6/vigia-go/internal/data/store/camera"
	"github.com/filipeandrade6/vigia-go/internal/sys/auth"

	"github.com/golang-migrate/migrate/v4"
	"go.uber.org/zap"
)

// TODO registrar no log os erros

type GerenciaService struct {
	pb.UnimplementedGerenciaServer
	log         *zap.SugaredLogger
	auth        *auth.Auth
	cameraStore camera.Store
}

func NewGerenciaService(log *zap.SugaredLogger, auth *auth.Auth, cameraStore camera.Store) *GerenciaService {
	return &GerenciaService{
		log:         log,
		auth:        auth,
		cameraStore: cameraStore,
	}
}

func (g *GerenciaService) Migrate(ctx context.Context, req *pb.MigrateReq) (*pb.MigrateRes, error) {

	// TODO add claims/auth
	version := req.GetVersao()

	if err := migration.Migrate(ctx, version); err != nil {
		if errors.As(err, &migrate.ErrNoChange) {
			g.log.Infow("no change in migration")
		} else {
			g.log.Errorw("migrate", err)
		}
	}

	g.log.Infow(fmt.Sprintf("migrate to version %d", version))

	return nil, nil
}

func (g *GerenciaService) CreateCamera(ctx context.Context, req *pb.CreateCameraReq) (*pb.CreateCameraRes, error) {
	cam := camera.FromProto(req.Camera)

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		g.log.Errorw("auth", "claims missing from context")
		return nil, errors.New("claims missing from context")
	}

	camID, err := g.cameraStore.Create(ctx, claims, cam)
	if err != nil {
		g.log.Errorw("create camera", err)
		return nil, fmt.Errorf("create: %w", err)
	}

	return &pb.CreateCameraRes{CameraId: camID}, nil
}

func (g *GerenciaService) ReadCamera(ctx context.Context, req *pb.ReadCameraReq) (*pb.ReadCameraRes, error) {

	cam, err := g.cameraStore.QueryByID(ctx, req.GetCameraId())
	if err != nil {
		g.log.Errorw("query camera", err)
		return nil, fmt.Errorf("query: %w", err)
	}

	return &pb.ReadCameraRes{Camera: cam.ToProto()}, err
}

func (g *GerenciaService) ReadCameras(ctx context.Context, req *pb.ReadCamerasReq) (*pb.ReadCamerasRes, error) {

	query := req.GetQuery()
	pageNumber := int(req.GetPageNumber())
	rowsPerPage := int(req.GetRowsPerPage())

	cameras, err := g.cameraStore.Query(ctx, query, pageNumber, rowsPerPage)
	if err != nil {
		g.log.Errorw("query cameras", err)
		return nil, fmt.Errorf("query: %w", err)
	}

	return &pb.ReadCamerasRes{Cameras: cameras.ToProto()}, nil
}

func (g *GerenciaService) UpdateCamera(ctx context.Context, req *pb.UpdateCameraReq) (*pb.UpdateCameraRes, error) {

	cam := camera.FromProto(req.Camera)

	if err := g.cameraStore.Update(ctx, cam); err != nil {
		g.log.Errorw("update camera", err)
		return nil, fmt.Errorf("update: %w", err)
	}
	return nil, nil
}

func (g *GerenciaService) DeleteCamera(ctx context.Context, req *pb.DeleteCameraReq) (*pb.DeleteCameraRes, error) {

	if err := g.cameraStore.Delete(ctx, req.GetCameraId()); err != nil {
		g.log.Errorw("delete camera", err)
		return nil, fmt.Errorf("delete: %w", err)
	}

	return nil, nil
}
