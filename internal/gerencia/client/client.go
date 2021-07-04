// Client do serviço de gravação
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

	conn, err := grpc.Dial("localhost:12346", opts...)
	if err != nil {
		fmt.Println("Erro aqui no client")
		panic(err)
	}
	defer conn.Close()
	client := pb.NewGravacaoClient(conn)

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		teste(client, &pb.InfoServidorGravacaoParams{})
	}
}

func teste(client pb.GravacaoClient, cfg *pb.InfoServidorGravacaoParams) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	infoServidorGravacao, err := client.InfoServidor(ctx, cfg)
	if err != nil {
		fmt.Println("Erro na chamado da função InfoServidor no client")
		panic(err)
	}

	for i, proc := range infoServidorGravacao.Processos {
		fmt.Println(
			i,
			proc.GetCameraId(),
			proc.GetProcessadorCaminho(),
			proc.GetStatusProcesso(),
		)
	}
}
