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
		Host:         viper.GetString("database.host"),
		Port:         int32(viper.GetInt("database.port")),
		User:         viper.GetString("database.user"),
		Password:     viper.GetString("database.pass"),
		Dbname:       viper.GetString("database.name"),
		Poolmaxconns: int32(viper.GetInt("database.poolmaxconns")),
	}, nil
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
