package server

import (
	"context"
	"fmt"
	"net"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/gerencia/models"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type gerenciaServer struct {
	pb.UnimplementedGerenciaServer
}

func (s *gerenciaServer) RegistrarServidorGravacao(ctx context.Context, req *pb.RegistrarServidorGravacaoReq) (*pb.RegistrarServidorGravacaoResp, error) {
	sg := &models.ServidorGravacao{}
	sg.FromProtobuf(req)
	// TODO tratar a informação
	sg.ID = "1"
	sg.Status = "OK"
	return sg.ToProtobuf(), nil
}

func (s *gerenciaServer) ConfigBancoDeDados(ctx context.Context, req *pb.ConfigBancoDeDadosReq) (*pb.ConfigBancoDeDadosResp, error) {
	return &pb.ConfigBancoDeDadosResp{
		Host:         viper.GetString("DB_HOST"),
		Port:         int32(viper.GetInt("DB_PORT")),
		User:         viper.GetString("DB_USER"),
		Password:     viper.GetString("DB_PASS"),
		Dbname:       viper.GetString("DB_NAME"),
		Poolmaxconns: int32(viper.GetInt("DB_POOL")),
	}, nil
}

func NovoServidorGerencia() *grpc.Server {
	lis, err := net.Listen(
		viper.GetString("SERVER_CONN"),
		fmt.Sprintf(
			"%s:%d",
			viper.GetString("SERVER_ENDERECO"),
			viper.GetInt("SERVER_PORTA"),
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
