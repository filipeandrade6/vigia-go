package grpc

import (
	"context"
	"fmt"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/core/camera"
	"github.com/filipeandrade6/vigia-go/internal/core/processo"
	"github.com/filipeandrade6/vigia-go/internal/core/registro"
	"github.com/filipeandrade6/vigia-go/internal/core/veiculo"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/service/processador"
	"github.com/filipeandrade6/vigia-go/internal/sys/database"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GravacaoService struct {
	pb.UnimplementedGravacaoServer

	log *zap.SugaredLogger

	servidorGravacaoID string
	armazenamento      string
	horasRetencao      int

	gerencia *GerenciaClient

	cameraCore   camera.Core
	processoCore processo.Core
	registroCore registro.Core
	veiculoCore  veiculo.Core

	processador *processador.Processador
	errChan     chan error
	matchChan   chan string
}

func NewGravacaoService(log *zap.SugaredLogger, armazenamento string, horasRetencao int) *GravacaoService {
	return &GravacaoService{
		log:           log,
		armazenamento: armazenamento,
		horasRetencao: horasRetencao,
		errChan:       make(chan error),
		matchChan:     make(chan string),
	}
}

// TODO colocar de modo atomico - db e registro
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
		return &pb.RegistrarRes{}, status.Error(codes.Internal, fmt.Sprintf("could not connect open database: %s", err))
	}

	gerenciaClient, err := NewClientGerencia(req.GetEnderecoIp(), int(req.GetPorta()))
	if err != nil {
		return &pb.RegistrarRes{}, status.Error(codes.Internal, fmt.Sprintf("could not connect to gRPC server: %s", err))
	}
	g.gerencia = gerenciaClient

	g.servidorGravacaoID = req.GetServidorGravacaoId()

	g.cameraCore = camera.NewCore(g.log, db)
	g.processoCore = processo.NewCore(g.log, db)
	g.registroCore = registro.NewCore(g.log, db)
	g.veiculoCore = veiculo.NewCore(g.log, db)

	g.processador = processador.New(
		g.registroCore,
		g.servidorGravacaoID,
		g.armazenamento,
		g.horasRetencao,
		g.errChan,
		g.matchChan,
	)

	go func() {
		for {
			select {
			case err := <-g.errChan:
				fmt.Println(err)

			case m := <-g.matchChan:
				fmt.Println(m)

			}
		}
	}()

	g.processador.Start()

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

func (g *GravacaoService) StartProcessos(ctx context.Context, req *pb.StartProcessosReq) (*pb.StartProcessosRes, error) {
	prcs := req.GetProcessos()

	for _, prc := range prcs {
		p, err := g.processoCore.QueryByID(ctx, prc)
		if err != nil {
			return &pb.StartProcessosRes{}, status.Error(codes.Internal, fmt.Sprintf("query database: %w", err))
		}
	}

	return &pb.StartProcessosRes{}, nil
}

func (g *GravacaoService) StopProcessos(ctx context.Context, req *pb.StopProcessosReq) (*pb.StopProcessosRes, error) {
	return &pb.StopProcessosRes{}, nil
}

func (g *GravacaoService) ListProcessos(ctx context.Context, req *pb.ListProcessosReq) (*pb.ListProcessosRes, error) {
	return &pb.ListProcessosRes{}, nil
}

func (g *GravacaoService) AtualizarMatchlist(ctx context.Context, req *pb.AtualizarMatchlistReq) (*pb.AtualizarMatchlistRes, error) {
	return &pb.AtualizarMatchlistRes{}, nil
}
func (g *GravacaoService) AtualizarHousekeeper(ctx context.Context, req *pb.AtualizarHousekeeperReq) (*pb.AtualizarHousekeeperRes, error) {
	return &pb.AtualizarHousekeeperRes{}, nil
}
func (g *GravacaoService) StartHousekeeper(ctx context.Context, req *pb.StartHousekeeperReq) (*pb.StartHousekeeperRes, error) {
	return &pb.StartHousekeeperRes{}, nil
}
func (g *GravacaoService) StopHousekeeper(ctx context.Context, req *pb.StopHousekeeperReq) (*pb.StopHousekeeperRes, error) {
	return &pb.StopHousekeeperRes{}, nil
}
func (g *GravacaoService) StatusHousekeeper(ctx context.Context, req *pb.StatusHousekeeperReq) (*pb.StatusHousekeeperRes, error) {
	return &pb.StatusHousekeeperRes{}, nil
}
func (g *GravacaoService) GetServidorInfo(ctx context.Context, req *pb.GetServidorInfoReq) (*pb.GetServidorInfoRes, error) {
	return &pb.GetServidorInfoRes{}, nil
}
