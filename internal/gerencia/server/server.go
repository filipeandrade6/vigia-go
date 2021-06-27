package server

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/filipeandrade6/vigia-go/internal/pb"
	"google.golang.org/grpc"
)

type gerenciaServer struct {
	pb.UnimplementedGerenciaServer
	mu sync.Mutex

	dbCfg *pb.DatabaseConfig
}

func (s *gerenciaServer) GetDatabaseConfig(ctx context.Context, cfg *pb.GravacaoConfig) (*pb.DatabaseConfig, error) {
	s.mu.Lock()
	fmt.Println("dentro do lock: ", cfg.Id)
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
	lis, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		fmt.Println("erro aqui")
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGerenciaServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
