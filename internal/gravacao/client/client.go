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
		dbUser:         db.User,
		dbPass:         db.Password,
		dbHost:         db.Host,
		dbPort:         int(db.Port),
		dbName:         db.DBName,
		dbPoolMaxConns: int(db.PoolMaxConns),
	}
}

func GetGerenciaClient(url string) *GerenciaClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(url, opts...)
	if err != nil {
		// TODO mudar isso aqui
		fmt.Println("Erro aqui no client")
		panic(err)
	}
	defer conn.Close()

	return &GerenciaClient{
		c: pb.NewGerenciaClient(conn),
	}
}
