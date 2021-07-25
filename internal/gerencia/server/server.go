package server

import (
	"context"
	"fmt"
	"net"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type gerenciaServer struct {
	pb.UnimplementedGerenciaServer
}

func (s *gerenciaServer) RegistrarServidorGravacao(ctx context.Context, req *pb.RegistrarServidorGravacaoReq) (*pb.RegistrarServidorGravacaoResp, error) {
	fmt.Println(req)

	return &pb.RegistrarServidorGravacaoResp{}, nil
}

func (s *gerenciaServer) ConfigBancoDeDados(ctx context.Context, req *pb.ConfigBancoDeDadosReq) (*pb.ConfigBancoDeDadosResp, error) {
	fmt.Println(req)

	return &pb.ConfigBancoDeDadosResp{}, nil
}

func NovoServidorGerencia() *grpc.Server {
	lis, err := net.Listen(
		viper.GetString("server.conn"),
		fmt.Sprintf(
			"%s:%d",
			viper.GetString("server.endereco"),
			viper.GetInt("server.porta"),
		),
	) // e.g. "tcp", "localhost:12346"
	if err != nil {
		panic(err)
	}

	grpcGerenciaServer := grpc.NewServer()
	pb.RegisterGerenciaServer(grpcGerenciaServer, &gerenciaServer{})
	grpcGerenciaServer.Serve(lis)

	return grpcGerenciaServer
}
