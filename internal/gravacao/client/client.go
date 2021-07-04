// Client do serviço de gerencia
package client

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"google.golang.org/grpc"
)

func Main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial("localhost:12347", opts...)
	if err != nil {
		fmt.Println("Erro aqui no client")
		panic(err)
	}
	defer conn.Close()
	client := pb.NewGerenciaClient(conn)

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		teste(
			client,
			&pb.GravacaoConfigReq{
				ServidorGravacao: "localhost",
			},
		)
	}
}

func teste(client pb.GerenciaClient, cfg *pb.GravacaoConfigReq) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	GravacaoConfigResp, err := client.GravacaoConfig(ctx, cfg)
	if err != nil {
		fmt.Println("Erro na chamado da função InfoServidor no client")
		panic(err)
	}

	fmt.Println(
		GravacaoConfigResp.GetServidorGravacaoId(),
		GravacaoConfigResp.GetDatabase(),
	)
}
