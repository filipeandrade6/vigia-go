package grpc

import (
	"context"
	"fmt"
	"time"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"

	"google.golang.org/grpc"
)

type GerenciaClient struct {
	c pb.GerenciaClient
}

func NewClientGerencia(endereco_ip string, porta int) (*GerenciaClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure()) // TODO trocar para SSL/TLS
	opts = append(opts, grpc.WithBlock())

	dialAddr := fmt.Sprintf("%s:%d", endereco_ip, porta)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, dialAddr, opts...)
	if err != nil {
		return nil, err
	}
	// conn, err := grpc.Dial(dialAddr, opts...)

	return &GerenciaClient{
		c: pb.NewGerenciaClient(conn),
	}, nil
}
