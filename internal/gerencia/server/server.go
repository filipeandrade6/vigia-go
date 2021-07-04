package server

import (
	"context"
	"fmt"
	"net"
	"sync"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"google.golang.org/grpc"
)

type gerenciaServer struct {
	pb.UnimplementedGerenciaServer
	mu sync.Mutex
}

func newServer() *gerenciaServer {
	return &gerenciaServer{}
}

func (s *gerenciaServer) GravacaoConfig(ctx context.Context, in *pb.GravacaoConfigReq) (*pb.GravacaoConfigResp, error) {
	s.mu.Lock()

	defer s.mu.Unlock()
	fmt.Println("Entrouaa qui")
	return &pb.GravacaoConfigResp{}, nil
}

func Main() {
	lis, err := net.Listen("tcp", "localhost:12347")
	if err != nil {
		fmt.Println("Erro aqui2")
		panic(err)
	}
	grpcGerenciaServer := grpc.NewServer()
	pb.RegisterGerenciaServer(grpcGerenciaServer, newServer())
	grpcGerenciaServer.Serve(lis)
}
