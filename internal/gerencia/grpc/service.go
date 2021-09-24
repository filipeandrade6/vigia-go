package grpc

import (
	"context"
	"fmt"

	pb "github.com/filipeandrade6/vigia-go/internal/api"
	"github.com/filipeandrade6/vigia-go/internal/data/migration"
	"github.com/filipeandrade6/vigia-go/internal/data/store/camera"
	gerenciaService "github.com/filipeandrade6/vigia-go/internal/gerencia/service"
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

func (g *gerenciaGRPCService) CreateServidorGravacao(ctx context.Context, req *pb.CreateServidorGravacaoReq) (*pb.CreateServidorGravacaoRes, error) {
	return &pb.CreateServidorGravacaoRes{}, nil
}

func (g *gerenciaGRPCService) CreateCamera(ctx context.Context, req *pb.CreateCameraReq) (*pb.CreateCameraRes, error) {

	cam := camera.FromProto(req.Camera)

	camID, err := g.gerenciaService.CreateCamera(ctx, cam)
	if err != nil {
		return &pb.CreateCameraRes{}, err
	}

	return &pb.CreateCameraRes{Camera: &pb.Camera{CameraId: camID}}, nil
}

func (g *gerenciaGRPCService) CreateProcesso(ctx context.Context, req *pb.CreateProcessoReq) (*pb.CreateProcessoRes, error) {
	return &pb.CreateProcessoRes{ProcessoId: "07c96eee-ab2f-4c17-b345-5634de4e2aac"}, nil
}

func (g *gerenciaGRPCService) Migrate(ctx context.Context, req *pb.MigrateReq) (*pb.MigrateRes, error) {
	fmt.Println(req.GetVersao())
	if err := migration.Migrate(context.Background()); err != nil {
		g.log.Fatalw("failed to migrate", err)
	}
	return &pb.MigrateRes{}, nil
}
