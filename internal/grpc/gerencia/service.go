package gerencia

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/data/migration"
	"github.com/filipeandrade6/vigia-go/internal/data/store/camera"
	gerenciaService "github.com/filipeandrade6/vigia-go/internal/gerencia/service"
	"github.com/filipeandrade6/vigia-go/internal/grpc/gerencia/pb"
	"github.com/filipeandrade6/vigia-go/internal/sys/auth"
	"github.com/golang-migrate/migrate/v4"
	"go.uber.org/zap"
)

// TODO posso ignorar o service e passar tudo para esse aqui
// TODO ou posso colocar mais uma camada estilo core do ardanlabs/service
type gerenciaGRPCService struct {
	pb.UnimplementedGerenciaServer // TODO remover isso aqui se eu implementar todos os serviços
	log                            *zap.SugaredLogger
	gerenciaService                *gerenciaService.GerenciaService
	// validator
}

// TODO ver o necessidade do ponteiro gerenciaService
func NewGerenciaService(log *zap.SugaredLogger, gerenciaService *gerenciaService.GerenciaService) *gerenciaGRPCService {
	return &gerenciaGRPCService{
		log:             log,
		gerenciaService: gerenciaService,
	}
}

func (g *gerenciaGRPCService) Migrate(ctx context.Context, req *pb.MigrateReq) (*pb.MigrateRes, error) {

	fmt.Println(req.GetVersao()) // TODO remover isso aqui

	if err := migration.Migrate(context.Background()); err != nil {
		if errors.As(err, &migrate.ErrNoChange) {
			g.log.Infow("no change in migration")
		} else {
			g.log.Fatalf("failed to migrate", err)
		}
	}
	return &pb.MigrateRes{}, nil
}

func (g *gerenciaGRPCService) CreateCamera(ctx context.Context, req *pb.CreateCameraReq) (*pb.CreateCameraRes, error) {

	cam := camera.FromProto(req.Camera)

	now := time.Now() // TODO pegar do contexto?

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return &pb.CreateCameraRes{}, errors.New("claims missing from context")
	}

	camID, err := g.gerenciaService.CreateCamera(ctx, claims, cam, now)
	if err != nil {
		return &pb.CreateCameraRes{}, err
	}

	return &pb.CreateCameraRes{Camera: &pb.Camera{CameraId: camID}}, nil
}

func (g *gerenciaGRPCService) ReadCamera(ctx context.Context, req *pb.ReadCameraReq) (*pb.ReadCameraRes, error) {

	cam := camera.FromProto(req.Camera)

	c, err := g.gerenciaService.ReadCamera(ctx, cam.CameraID)
	if err != nil {
		return &pb.ReadCameraRes{}, err
	}

	return &pb.ReadCameraRes{Camera: c.ToProto()}, err
}

// TODO colocar pageNumer e rowsPerPage no proto
func (g *gerenciaGRPCService) ReadCameras(ctx context.Context, req *pb.ReadCamerasReq) (*pb.ReadCamerasRes, error) {

	c, err := g.gerenciaService.ReadCameras(ctx, 1, 1000)
	if err != nil {
		return &pb.ReadCamerasRes{}, err
	}

	return &pb.ReadCamerasRes{Cameras: c.ToProto()}, nil
}

func (g *gerenciaGRPCService) UpdateCamera(ctx context.Context, req *pb.UpdateCameraReq) (*pb.UpdateCameraRes, error) {

	now := time.Now()

	cam := camera.FromProto(req.Camera)

	if err := g.gerenciaService.UpdateCamera(ctx, cam, now); err != nil {
		return &pb.UpdateCameraRes{}, err
	}

	return &pb.UpdateCameraRes{}, nil
}

func (g *gerenciaGRPCService) DeleteCamera(ctx context.Context, req *pb.DeleteCameraReq) (*pb.DeleteCameraRes, error) {

	cam := camera.FromProto(req.Camera)

	if err := g.gerenciaService.DeleteCamera(ctx, cam.CameraID); err != nil {
		return &pb.DeleteCameraRes{}, err
	}

	return &pb.DeleteCameraRes{}, nil
}
