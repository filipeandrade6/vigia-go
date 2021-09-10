// Server de serviço de gravação
package server

// verificar health in features no exemplo do gRPC

import (
	"context"
	"fmt"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"google.golang.org/grpc"
)

type GravacaoServer struct {
	pb.UnimplementedGravacaoServer
}

func NovoServidorGravacao() *grpc.Server {
	grpcGravacaoServer := grpc.NewServer()
	pb.RegisterGravacaoServer(grpcGravacaoServer, &GravacaoServer{})

	return grpcGravacaoServer
}

func (s *GravacaoServer) IniciarProcessamento(ctx context.Context, req *pb.IniciarProcessamentoReq) (*pb.IniciarProcessamentoResp, error) {
	return &pb.IniciarProcessamentoResp{
		Status: "chegou",
	}, nil
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
