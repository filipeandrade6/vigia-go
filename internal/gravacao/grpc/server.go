package grpc

import (
	"context"
	"fmt"
	"os"

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

	gerencia *GerenciaClient

	cameraCore   camera.Core
	processoCore processo.Core
	registroCore registro.Core
	veiculoCore  veiculo.Core

	processador *processador.Processador
	errChan     chan error
	matchChan   chan string
}

func NewGravacaoService(log *zap.SugaredLogger) *GravacaoService {
	return &GravacaoService{
		log:       log,
		errChan:   make(chan error),
		matchChan: make(chan string),
	}
}

func (g *GravacaoService) Registrar(ctx context.Context, req *pb.RegistrarReq) (*pb.RegistrarRes, error) {
	if g.gerencia != nil {
		return &pb.RegistrarRes{}, status.Error(codes.AlreadyExists, "ja possui servidor de gerencia registrado")
	}

	g.log.Infow("start", "teste", "teste")

	err := os.MkdirAll(req.GetArmazenamento(), os.ModePerm)
	if err != nil {
		return &pb.RegistrarRes{}, status.Error(codes.InvalidArgument, err.Error())
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

	// TODO o DB quando a altenticação falha com senha diferente ele não alerta

	gerenciaClient, err := NewClientGerencia(req.GetEnderecoIp(), int(req.GetPorta()))
	if err != nil {
		return &pb.RegistrarRes{}, status.Error(codes.Internal, fmt.Sprintf("could not connect to gRPC server: %s", err))
	}
	g.gerencia = gerenciaClient

	g.cameraCore = camera.NewCore(g.log, db)
	g.processoCore = processo.NewCore(g.log, db)
	g.registroCore = registro.NewCore(g.log, db)
	g.veiculoCore = veiculo.NewCore(g.log, db)

	g.processador = processador.New(req.GetServidorGravacaoId(), req.GetArmazenamento(), int(req.GetHorasRetencao()), g.registroCore, g.errChan, g.matchChan)

	go g.start()
	go g.processador.Start()

	return &pb.RegistrarRes{}, nil
}

func (g *GravacaoService) start() {
	for {
		select {
		case err := <-g.errChan:
			fmt.Println(err)

		case m := <-g.matchChan:
			fmt.Println(m)
		}
	}
}

func (g *GravacaoService) RemoverRegistro(ctx context.Context, req *pb.RemoverRegistroReq) (*pb.RemoverRegistroRes, error) {
	if g.gerencia == nil {
		return &pb.RemoverRegistroRes{}, status.Error(codes.NotFound, "nao possui servidor de gerencia registrado")
	}
	g.gerencia = nil

	g.processador.Stop()

	// TODO interromper os cores, processador, e db?

	return &pb.RemoverRegistroRes{}, nil
}

func (g *GravacaoService) StartProcessos(ctx context.Context, req *pb.StartProcessosReq) (*pb.StartProcessosRes, error) {
	prcs := req.GetProcessos()

	var processos []processador.Processo
	for _, prc := range prcs {
		p, err := g.processoCore.QueryByID(ctx, prc)
		if err != nil {
			return &pb.StartProcessosRes{}, status.Error(codes.Internal, fmt.Sprintf("query database: %s", err))
		}

		c, err := g.cameraCore.QueryByID(ctx, p.CameraID)
		if err != nil {
			return &pb.StartProcessosRes{}, status.Error(codes.Internal, fmt.Sprintf("query database: %s", err))
		}

		processos = append(processos, processador.Processo{
			ProcessoID:  prc,
			EnderecoIP:  c.EnderecoIP,
			Porta:       c.Porta,
			Canal:       c.Canal,
			Usuario:     c.Usuario,
			Senha:       c.Senha,
			Processador: p.Processador,
		})
	}

	g.processador.StartProcessos(processos)

	return &pb.StartProcessosRes{}, nil
}

func (g *GravacaoService) StopProcessos(ctx context.Context, req *pb.StopProcessosReq) (*pb.StopProcessosRes, error) {
	err := g.processador.StopProcessos(req.GetProcessos())
	if err != nil {
		return &pb.StopProcessosRes{}, status.Error(codes.Internal, fmt.Sprintf("stopping processo: %s", err))
	}

	return &pb.StopProcessosRes{}, nil
}

func (g *GravacaoService) ListProcessos(ctx context.Context, req *pb.ListProcessosReq) (*pb.ListProcessosRes, error) {
	processos, retry := g.processador.ListProcessos()

	return &pb.ListProcessosRes{ProcessosEmExecucao: processos, ProcessosEmTentativa: retry}, nil
}

func (g *GravacaoService) UpdateMatchlist(ctx context.Context, req *pb.UpdateMatchlistReq) (*pb.UpdateMatchlistRes, error) {

	veiculos, err := g.veiculoCore.QueryAll(ctx)
	if err != nil {
		return &pb.UpdateMatchlistRes{}, status.Error(codes.Internal, fmt.Sprintf("query database: %s", err))
	}

	var matchlist []string
	for _, v := range veiculos {
		matchlist = append(matchlist, v.Placa)
	}

	g.processador.UpdateMatchlist(matchlist)

	return &pb.UpdateMatchlistRes{}, nil
}

func (g *GravacaoService) UpdateArmazenamento(ctx context.Context, req *pb.UpdateArmazenamentoReq) (*pb.UpdateArmazenamentoRes, error) {
	err := g.processador.UpdateArmazenamento(req.GetArmazenamento(), int(req.GetHorasRetencao()))
	if err != nil {
		return &pb.UpdateArmazenamentoRes{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.UpdateArmazenamentoRes{}, nil
}
