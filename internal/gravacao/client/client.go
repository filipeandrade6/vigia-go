// Client do serviço de gerencia
package client

import (
	"fmt"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	// "github.com/filipeandrade6/vigia-go/internal/database"
	// "github.com/filipeandrade6/vigia-go/internal/gravacao/models"
	"github.com/spf13/viper"

	"google.golang.org/grpc"
)

type GerenciaClient struct {
	c pb.GerenciaClient
}

// func (g *GerenciaClient) ConfigBancoDeDados() *database.Config {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()

// 	configBancoDeDadosResp, err := g.c.ConfigBancoDeDados(ctx, nil)
// 	if err != nil {
// 		fmt.Println("Erro na chamada da função GravacaoConfig no client")
// 		panic(err)
// 	}

// 	var db *models.Database
// 	db.FromProtobuf(configBancoDeDadosResp)

// 	return &database.Config{
// 		DBUser:         db.User,
// 		DBPass:         db.Password,
// 		DBHost:         db.Host,
// 		DBPort:         int(db.Port),
// 		DBName:         db.DBName,
// 		DBPoolMaxConns: int(db.PoolMaxConns),
// 	}
// }

func NovoClientGerencia() *GerenciaClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(
		fmt.Sprintf(
			"%s:%d",
			viper.GetString("CLIENT_ENDERECO"),
			viper.GetInt("CLIENT_PORTA"),
		), opts...)
	if err != nil {
		fmt.Println("Erro aqui no client") // TODO mudar isso aqui
		panic(err)
	}
	defer conn.Close()

	return &GerenciaClient{
		c: pb.NewGerenciaClient(conn),
	}
}
