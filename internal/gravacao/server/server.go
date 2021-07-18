// Server de serviço de gravação
package server

// verificar health in features no exemplo do gRPC

import (
	"context"
	"fmt"
	"net"
	"sync"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"google.golang.org/grpc"
)

type gravacaoServer struct {
	pb.UnimplementedGravacaoServer
	mu sync.Mutex
}

func novoServidor() *gravacaoServer {
	return &gravacaoServer{}
}

func (s *gravacaoServer) InfoServidor(ctx context.Context, in *pb.InfoServidorReq) (*pb.InfoServidorResp, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Println("Entrou aqui")
	return &pb.InfoServidorResp{
		Processos: []*pb.InfoServidorResp_Processo{
			{
				CameraId:           1,
				ProcessadorCaminho: "/usr/local/vigia-go/processadores/proc1",
				StatusProcesso:     pb.InfoServidorResp_Processo_PARADO,
			},
		},
	}, nil
}

func StartServer(tipo, url string) {
	lis, err := net.Listen(tipo, url)
	if err != nil {
		fmt.Println("Erro aqui")
		panic(err)
	}
	grpcGravacaoServer := grpc.NewServer()
	pb.RegisterGravacaoServer(grpcGravacaoServer, novoServidor())
	grpcGravacaoServer.Serve(lis)
}
