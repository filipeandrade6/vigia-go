// Client do serviço de gravação
package client

import (
	"context"
	"fmt"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GravacaoClient struct {
	c pb.GravacaoClient
}

func (g *GravacaoClient) IniciarProcessamento(context.Context, *pb.IniciarProcessamentoReq) (*pb.IniciarProcessamentoResp, error) {
	fmt.Println("acessou fund IniciarProcesamento")
	return nil, status.Errorf(codes.Unimplemented, "method IniciarProcessamento not implemented")
}

func NovoClientGravacao() *GravacaoClient {
	fmt.Println("chegou no novoclientGravacao")

	conn, err := grpc.Dial(
		fmt.Sprintf(
			"%s:%d",
			viper.GetString("GRA_CLIENT_ENDERECO"),
			viper.GetInt("GRA_CLIENT_PORTA"),
		), grpc.WithInsecure())
	if err != nil {
		fmt.Println("Erro aqui no client") // TODO mudar isso aqui
		panic(err)
	}
	defer conn.Close()

	fmt.Println("quase no final do novoClient")

	return &GravacaoClient{
		c: pb.NewGravacaoClient(conn),
	}
}
