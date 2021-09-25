package gravacao

import (
	"context"

	gravacaoService "github.com/filipeandrade6/vigia-go/internal/gravacao/service"
	"github.com/filipeandrade6/vigia-go/internal/grpc/gravacao/pb"
	"go.uber.org/zap"
)

type gravacaoGRPCService struct {
	pb.UnimplementedGravacaoServer
	log             *zap.SugaredLogger
	gravacaoService *gravacaoService.GravacaoService
}

// TODO ver a necessidade do ponteiro gravacaoService
func NewGravacaoService(log *zap.SugaredLogger, gravacaoService *gravacaoService.GravacaoService) *gravacaoGRPCService {
	return &gravacaoGRPCService{
		log:             log,
		gravacaoService: gravacaoService,
	}
}

func (g *gravacaoGRPCService) IniciarProcessamento(ctx context.Context, req *pb.IniciarProcessamentoReq) (*pb.IniciarProcessamentoResp, error) {
	return &pb.IniciarProcessamentoResp{Status: "ok"}, nil
}
