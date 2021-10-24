package grpc

import (
	"context"
	"fmt"
	"time"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/sys/operrors"

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

func (g *GerenciaClient) Match(registroID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	r := &pb.MatchReq{RegistroId: registroID}

	if _, err := g.c.Match(ctx, r); err != nil {
		return err
	}

	return nil
}

func (g *GerenciaClient) ErrorReport(err operrors.OpError) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	r := &pb.ErrorReportReq{
		ServidorGravacaoId: err.ServidorGravacaoID,
		ProcessoId:         err.ProcessoID,
		RegistroId:         err.RegistroID,
		Error:              err.Err.Error(),
		StoppedProcesso:    err.StoppedProcesso,
	}

	if _, err := g.c.ErrorReport(ctx, r); err != nil {
		return err
	}

	return nil
}
