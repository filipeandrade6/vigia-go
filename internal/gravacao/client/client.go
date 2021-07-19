// Client do serviço de gerencia
package client

import (
	"context"
	"fmt"
	"time"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/database"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/models"

	"google.golang.org/grpc"
)

type GerenciaClient struct {
	c pb.GerenciaClient
}

func (g *GerenciaClient) GetDatabase() *database.Config {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	databaseConfigResp, err := g.c.DatabaseConfig(ctx, nil)
	if err != nil {
		fmt.Println("Erro na chamada da função GravacaoConfig no client")
		panic(err)
	}

	var db *models.Database
	db.FromProtobuf(databaseConfigResp)

	return &database.Config{
		DBUser:         db.User,
		DBPass:         db.Password,
		DBHost:         db.Host,
		DBPort:         int(db.Port),
		DBName:         db.DBName,
		DBPoolMaxConns: int(db.PoolMaxConns),
	}
}

func NovoClientGerencia(url string) *GerenciaClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(url, opts...)
	if err != nil {
		fmt.Println("Erro aqui no client") // TODO mudar isso aqui
		panic(err)
	}
	defer conn.Close()

	return &GerenciaClient{
		c: pb.NewGerenciaClient(conn),
	}
}
