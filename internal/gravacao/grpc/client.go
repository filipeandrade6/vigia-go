package grpc

import (
	"fmt"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"

	"google.golang.org/grpc"
)

type GerenciaClient struct {
	c pb.GerenciaClient
}

func NewClientGerencia(endereco_ip string, porta int) *GerenciaClient {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	dialAddr := fmt.Sprintf("%s:%d", endereco_ip, porta)

	conn, err := grpc.Dial(dialAddr, opts...)
	if err != nil {
		fmt.Println("Erro aqui no client") // TODO mudar isso aqui
		panic(err)
	}
	//defer conn.Close() // TODO esse aqui vai dar BO

	return &GerenciaClient{
		c: pb.NewGerenciaClient(conn),
	}
}

// TODO preciso registrar o HealthServer no GRPC
