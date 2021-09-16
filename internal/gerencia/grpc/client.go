package grpc

import (
	"context"
	"fmt"
	"time"

	pb "github.com/filipeandrade6/vigia-go/internal/api"
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
	fmt.Println("chegaste aqui em novoclientgerencia")

	cfg := fmt.Sprintf(
		"%s:%d",
		viper.GetString("GER_HOST"), // TODO assim como no DB juntar endereco e porta em uma unica var
		viper.GetInt("GER_SERVER_PORT"),
	)

	fmt.Println("as config de conexão é", cfg)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(cfg, opts...)
	if err != nil {
		fmt.Println("Erro aqui no client") // TODO mudar isso aqui
		panic(err)
	}
	//defer conn.Close() // TODO esse aqui vai dar BO

	fmt.Println("criado client de gerencia")

	return &GerenciaClient{
		c: pb.NewGerenciaClient(conn),
	}
}

func (g *GerenciaClient) Migrate() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if _, err := g.c.Migrate(ctx, &pb.MigrateReq{Versao: 5}); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (g *GerenciaClient) CreateCamera() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cam := &pb.CreateCameraReq{
		Descricao:      "Teste",
		EnderecoIp:     "10.0.0.1",
		Porta:          12,
		Canal:          1,
		Usuario:        "admin",
		Senha:          "admin",
		Geolocalizacao: "-12.3242, -45.1234",
	}

	camID, err := g.c.CreateCamera(ctx, cam)
	if err != nil {
		return err
	}

	fmt.Println(camID.GetCameraId())

	return nil
}
