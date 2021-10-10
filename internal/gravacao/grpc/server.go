package grpc

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

type Processo struct {
}

type GravacaoService struct {
	pb.UnimplementedGravacaoServer

	log *zap.SugaredLogger

	gerencia       *GerenciaClient
	processos      map[string]Processo
	processoErrors chan error

	registroCore registro.Core
	veiculoCore  veiculo.Core
}

func NewGravacaoService(log *zap.SugaredLogger) *GravacaoService {
	processoErrors := make(chan error)
	processos := make(map[string]Processo)

	return &GravacaoService{
		log:            log,
		processos:      processos,
		processoErrors: processoErrors,
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

	// TODO interromper os core?

	return &pb.RemoverRegistroRes{}, nil
}

func (g *GravacaoService) CreateProcesso(ctx context.Context, req *pb.CreateProcessoReq) (*pb.CreateProcessoRes, error) {
	return &pb.CreateProcessoRes{}, nil
}

func (g *GravacaoService) ReadProcesso(ctx context.Context, req *pb.ReadProcessoReq) (*pb.ReadProcessoRes, error) {
	return &pb.ReadProcessoRes{}, nil
}

func (g *GravacaoService) ReadProcessos(ctx context.Context, req *pb.ReadProcessosReq) (*pb.ReadProcessosRes, error) {
	return &pb.ReadProcessosRes{}, nil
}

func (g *GravacaoService) UpdateProcesso(ctx context.Context, req *pb.UpdateProcessoReq) (*pb.UpdateProcessoRes, error) {
	return &pb.UpdateProcessoRes{}, nil
}

func (g *GravacaoService) DeleteProcesso(ctx context.Context, req *pb.DeleteProcessoReq) (*pb.DeleteProcessoRes, error) {
	return &pb.DeleteProcessoRes{}, nil
}
