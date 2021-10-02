package service

import (
	"context"
	"fmt"
	"time"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/core/camera"
	"github.com/filipeandrade6/vigia-go/internal/core/usuario"
	"github.com/golang/protobuf/ptypes/wrappers"

	"google.golang.org/grpc"
)

type GerenciaClient struct {
	c pb.GerenciaClient
}

func NewClientGerencia(dialAddr string) *GerenciaClient {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(dialAddr, opts...)
	if err != nil {
		fmt.Println("Erro aqui no client") // TODO mudar isso aqui
		panic(err)
	}
	//defer conn.Close() // TODO esse aqui vai dar BO

	return &GerenciaClient{
		c: pb.NewGerenciaClient(conn),
	}
}

func (g *GerenciaClient) Migrate() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// TODO refatorar aqui
	if _, err := g.c.Migrate(ctx, &pb.MigrateReq{Versao: 7}); err != nil {
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

	return c.GetCameraId(), nil
}

func (g *GerenciaClient) ReadCamera(cameraID string) (camera.Camera, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c, err := g.c.ReadCamera(ctx, &pb.ReadCameraReq{CameraId: cameraID})
	if err != nil {
		return camera.Camera{}, err
	}

	return camera.FromProto(c.Camera), nil
}

func (g *GerenciaClient) ReadCameras(query string, pageNumber int, rowsPerPage int) (camera.Cameras, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := &pb.ReadCamerasReq{
		Query:       query,
		PageNumber:  int32(pageNumber),
		RowsPerPage: int32(rowsPerPage),
	}

	c, err := g.c.ReadCameras(ctx, request)
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

func (g *GerenciaClient) DeleteCamera(camerasID []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if _, err := g.c.DeleteCamera(ctx, &pb.DeleteCameraReq{CameraId: camerasID}); err != nil {
		return err
	}

	return nil
}

func (g *GerenciaClient) Login(email, senha string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	t, err := g.c.Login(ctx, &pb.LoginReq{Email: email, Senha: senha})
	if err != nil {
		return "", err
	}

	token := t.GetAccessToken()

	return token, nil
}

func (g *GerenciaClient) UpdateUsuario(usuarioID string, usuario usuario.UpdateUsuario) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if _, err := g.c.UpdateUsuario(ctx, &pb.UpdateUsuarioReq{
		UsuarioId: usuarioID,
		Funcao:    []string{"USER", "MANAGER"},
		Senha:     &wrappers.StringValue{Value: "secret"},
	}); err != nil {
		return err
	}

	return nil
}
