// Client do serviço de gravação
package client

import (
	"fmt"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type GravacaoClient struct {
	c pb.GravacaoClient
}

// func (g *GravacaoClient) ConfigurarProcesso(req *pb.ConfigurarProcessoReq) *models. { // TODO arrumar aqui
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()

// 	infoServidorResp, err := g.c.InfoServidor(ctx, nil)
// 	if err != nil {
// 		fmt.Println("Erro na chamada do client")
// 		panic(err)
// 	}

// 	var sv *models.ServidorGravacao
// 	sv.FromProtobuf(infoServidorResp)

// 	return &sv
// }

func NovoClientGravacao() *GravacaoClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(
		fmt.Sprintf(
			"%s:%d",
			viper.GetString("client.endereco"),
			viper.GetInt("client.porta"),
		),
		opts...)
	if err != nil {
		fmt.Println("Erro aqui no client") // TODO mudar isso aqui
		panic(err)
	}
	defer conn.Close()

	return &GravacaoClient{
		c: pb.NewGravacaoClient(conn),
	}
}
