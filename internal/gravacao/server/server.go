// Server de serviço de gravação
package server

// verificar health in features no exemplo do gRPC

import (
	"context"
	"fmt"
	"net"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type GravacaoServer struct {
	pb.UnimplementedGravacaoServer
}

func (s *GravacaoServer) InfoProcessos(ctx context.Context, req *pb.InfoProcessosReq) (*pb.InfoProcessosResp, error) {
	return &pb.InfoProcessosResp{
		Processos: []*pb.InfoProcessosResp_Processo{
			{
				Id:                 10,
				CameraId:           10,
				ProcessadorCaminho: "processador_a1",
				Status:             0, // TODO ver como utilizar nome da variavel no lugar de numero
			},
		},
	}, nil
}

func (s *GravacaoServer) ConfigurarProcesso(ctx context.Context, req *pb.ConfigurarProcessoReq) (*pb.ConfigurarProcessoResp, error) {
	fmt.Println(req)
	return &pb.ConfigurarProcessoResp{
		Status: 0, // TODO ver como utilizar nome da variavel no lugadr de inteiro
	}, nil
}

func NovoServidorGravacao(tipo, url string) *grpc.Server {
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

	grpcGravacaoServer := grpc.NewServer()
	pb.RegisterGravacaoServer(grpcGravacaoServer, &GravacaoServer{})
	grpcGravacaoServer.Serve(lis)

	return grpcGravacaoServer
}
