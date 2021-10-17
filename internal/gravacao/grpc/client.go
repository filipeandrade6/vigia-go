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

func (g *GerenciaClient) Check(servidorID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	r := &pb.CheckReq{ServidorGravacaoId: servidorID}
	if _, err := g.c.Check(ctx, r); err != nil {
		return err
	}

	return nil
}

func (g *GerenciaClient) Match(veiculo_id, registro_id string) {
	ctx := context.Background()

	r := &pb.MatchReq{
		VeiculoId:  veiculo_id,
		RegistroId: registro_id,
	}

	if _, err := g.c.Match(ctx, r); err != nil {
		// TODO adicionar no error? - colocar buffer? caso o gerencia esteja offline
	}

}

// func (g *GerenciaClient) ProcessoError(servidorGravacaoID, processoID string, errType int) {
// 	ctx := context.Background()

// 	r := &pb.ProcessoErrorReq{
// 		ServidorGravacaoId: servidorGravacaoID,
// 		ProcessoId: processoID,
// 		Erro: ,
// 	}

// 	if _, err := g.c.ProcessoError(ctx, r); err != nil {
// 		// TODO adicionar error? - colocar buffer? caso o gerencia esteja offline
// 	}
// }
