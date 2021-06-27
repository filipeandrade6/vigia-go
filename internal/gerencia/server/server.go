package server

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/filipeandrade6/vigia-go/internal/pb"
	"google.golang.org/grpc"
)

var (
	port = 10000
)

type gerenciaServer struct {
	pb.UnimplementedGerenciaServer
	mu sync.Mutex

	dbCfg *pb.DatabaseConfig
}

func (s *gerenciaServer) GetDatabaseConfig(ctx context.Context, in *pb.GravacaoConfig, opts ...grpc.CallOption) (*pb.DatabaseConfig, error) {
	s.mu.Lock()
	fmt.Println("dentro do lock: ", in.GetId())
	s.mu.Unlock()

	return s.dbCfg, nil
}

func newServer() *gerenciaServer {
	return &gerenciaServer{
		dbCfg: &pb.DatabaseConfig{
			Host:         "localhost",
			Port:         5432,
			User:         "postgres",
			Password:     "postgres",
			Dbname:       "vigia",
			Poolmaxconns: 50,
		},
	}
}

func Main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		fmt.Println("erro aqui")
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGerenciaServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
