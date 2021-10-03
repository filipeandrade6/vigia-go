package service

import (
	"context"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/core/registro"
	"github.com/filipeandrade6/vigia-go/internal/core/veiculo"
	"github.com/filipeandrade6/vigia-go/internal/sys/database"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GravacaoService struct {
	pb.UnimplementedGravacaoServer

	log *zap.SugaredLogger

	gerencia *GerenciaClient

	registroCore registro.Core
	veiculoCore  veiculo.Core
}

func NewGravacaoService(log *zap.SugaredLogger) *GravacaoService {
	return &GravacaoService{
		log: log,
	}
}

func (g *GravacaoService) Registrar(ctx context.Context, req *pb.RegistrarReq) (*pb.RegistrarRes, error) {
	if g.gerencia != nil {
		return &pb.RegistrarRes{}, status.Error(codes.AlreadyExists, "ja possui servidor de gerencia registrado")
	}

	db, err := database.Open(database.Config{
		User:         req.GetDbUser(),
		Password:     req.GetDbPassword(),
		Host:         req.GetDbHost(),
		Name:         req.GetDbName(),
		MaxIDLEConns: int(req.GetDbMaxidleconns()),
		MaxOpenConns: int(req.GetDbMaxopenconns()),
		SSLMode:      req.GetDbSslmode(),
	})
	if err != nil {
		return &pb.RegistrarRes{}, status.Error(codes.Internal, err.Error())
	}

	g.registroCore = registro.NewCore(g.log, db)
	g.veiculoCore = veiculo.NewCore(g.log, db)

	g.gerencia = NewClientGerencia(req.GetEnderecoIp(), int(req.GetPorta()))

	return &pb.RegistrarRes{}, nil
}

func (g *GravacaoService) RemoverRegistro(ctx context.Context, req *pb.RemoverRegistroReq) (*pb.RemoverRegistroRes, error) {
	if g.gerencia == nil {
		return &pb.RemoverRegistroRes{}, status.Error(codes.NotFound, "nao possui servidor de gerencia registrado")
	}
	g.gerencia = nil

	return &pb.RemoverRegistroRes{}, nil
}
