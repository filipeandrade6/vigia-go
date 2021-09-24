package grpc

import (
	"context"
	"fmt"
	"time"

	pb "github.com/filipeandrade6/vigia-go/internal/api"
	"github.com/filipeandrade6/vigia-go/internal/data/store/camera"

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
		viper.GetString("VIGIA_GER_HOST"), // TODO assim como no DB juntar endereco e porta em uma unica var
		viper.GetInt("VIGIA_GER_SERVER_PORT"),
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

	// TODO refatorar aqui
	if _, err := g.c.Migrate(ctx, &pb.MigrateReq{Versao: 5}); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (g *GerenciaClient) CreateCamera(cam camera.Camera) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c, err := g.c.CreateCamera(ctx, &pb.CreateCameraReq{Camera: cam.ToProto()})
	if err != nil {
		return "", err
	}

	camRes := camera.FromProto(c.Camera)

	return camRes.CameraID, nil
}
