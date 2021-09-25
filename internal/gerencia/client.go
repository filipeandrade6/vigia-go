package gravacao

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
	conn, err := grpc.Dial(
		fmt.Sprintf(
			"%s:%d",
			viper.GetString("VIGIA_GRA_HOST"),
			viper.GetInt("VIGIA_GRA_SERVER_PORT"),
		), grpc.WithInsecure())
	if err != nil {
		fmt.Println("Erro aqui no client") // TODO mudar isso aqui
		panic(err)
	}
	defer conn.Close()

	return &GravacaoClient{
		c: pb.NewGravacaoClient(conn),
	}
}
