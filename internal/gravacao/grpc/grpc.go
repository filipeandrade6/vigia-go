package grpc

import (
	"context"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/service"
	"go.uber.org/zap"
)

type gravacaoGRPCService struct {
	pb.UnimplementedGravacaoServer
	log             *zap.SugaredLogger
	gravacaoService service.GravacaoService
}

// TODO ver a necessidade do ponteiro gravacaoService
func NewGravacaoService(log *zap.SugaredLogger, gravacaoService *service.GravacaoService) *gravacaoGRPCService {
	return &gravacaoGRPCService{
		log:             log,
		gravacaoService: *gravacaoService,
	}
}

func (g *gravacaoGRPCService) IniciarProcessamento(ctx context.Context, req *pb.IniciarProcessamentoReq) (*pb.IniciarProcessamentoResp, error) {
	return &pb.IniciarProcessamentoResp{Status: "ok"}, nil
}
