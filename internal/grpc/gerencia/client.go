package gerencia

import (
	"context"
	"fmt"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/data/store/camera"
	"github.com/filipeandrade6/vigia-go/internal/grpc/gerencia/pb"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type GerenciaClient struct {
	c pb.GerenciaClient
}

func NewClientGerencia() *GerenciaClient {
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

func (g *GerenciaClient) ReadCamera(cameraID string) (camera.Camera, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c, err := g.c.ReadCamera(ctx, &pb.ReadCameraReq{Camera: camera.Camera{CameraID: cameraID}.ToProto()})
	if err != nil {
		return camera.Camera{}, err
	}

	return camera.FromProto(c.Camera), nil
}

func (g *GerenciaClient) ReadCameras() (camera.Cameras, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c, err := g.c.ReadCameras(ctx, &pb.ReadCamerasReq{})
	if err != nil {
		return camera.Cameras{}, err
	}

	cameras := camera.CamerasFromProto(c.Cameras)

	return cameras, nil
}

func (g *GerenciaClient) UpdateCamera(cam camera.Camera) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if _, err := g.c.UpdateCamera(ctx, &pb.UpdateCameraReq{Camera: cam.ToProto()}); err != nil {
		return err
	}

	return nil
}

func (g *GerenciaClient) DeleteCamera(cameraID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if _, err := g.c.DeleteCamera(ctx, &pb.DeleteCameraReq{Camera: camera.Camera{CameraID: cameraID}.ToProto()}); err != nil {
		return err
	}

	return nil
}