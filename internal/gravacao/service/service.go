package service

import (
	"fmt"
	"net"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/core/registro"
	"github.com/filipeandrade6/vigia-go/internal/core/veiculo"
	grpc_gravacao "github.com/filipeandrade6/vigia-go/internal/gravacao/grpc"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/service/processador"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GravacaoService struct {
	log *zap.SugaredLogger

	processador    *processador.Processador
	gerenciaClient *grpc_gravacao.GerenciaClient
	gravacaoServer *grpc.Server

	registroCore registro.Core
	veiculoCore  veiculo.Core

	sysErrors chan error
}

func NewGravacaoService(gravacaoServer *grpc.Server) *GravacaoService {
	return &GravacaoService{
		gravacaoServer: gravacaoServer,
	}
}

func (g *GravacaoService) Start() {
	// =========================================================================
	// gRPC Server
	sv := grpc_gravacao.NewGravacaoService(g.log)
	grpcServer := grpc.NewServer()
	pb.RegisterGravacaoServer(grpcServer, sv)
	go func() {
		lis, err := net.Listen(cfg.Gravacao.Conn, fmt.Sprintf(":%d", cfg.Gravacao.Port))
		if err != nil {
			g.log.Errorw("startup", "status", "could not open socket", cfg.Gravacao.Conn, cfg.Gravacao.Port, "ERROR", err)
		}

		g.log.Infow("startup", "status", "gRPC server started") // TODO adicionar host
		g.sysErrors <- grpcServer.Serve(lis)
	}()

	// =========================================================================
	// gRPC Server

}

func (g *GravacaoService) Stop() {
	g.gravacaoServer.GracefulStop()
	g.processador.Stop()
}
