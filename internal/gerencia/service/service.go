package service

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/data/migration"
	"github.com/filipeandrade6/vigia-go/internal/data/store/camera"
	"github.com/filipeandrade6/vigia-go/internal/data/store/usuario"
	"github.com/filipeandrade6/vigia-go/internal/sys/auth"
	"github.com/filipeandrade6/vigia-go/internal/sys/database"

	"github.com/golang-migrate/migrate/v4"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO registrar no log os erros

type GerenciaService struct {
	pb.UnimplementedGerenciaServer
	log          *zap.SugaredLogger
	auth         *auth.Auth
	cameraStore  camera.Store
	usuarioStore usuario.Store
}

func NewGerenciaService(log *zap.SugaredLogger, auth *auth.Auth, cameraStore camera.Store, usuarioStore usuario.Store) *GerenciaService {
	return &GerenciaService{
		log:          log,
		auth:         auth,
		cameraStore:  cameraStore,
		usuarioStore: usuarioStore,
	}
}

func (g *GerenciaService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	// TODO se o servico chamado for Login - so encaminha o contexto
	fmt.Println(fullMethodName)
	if fullMethodName == "/gerencia.Gerencia/Login" || fullMethodName == "/gerencia.Gerencia/Migrate" {
		return ctx, nil
	}

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token")
	}
	fmt.Println(token)

	claims, err := g.auth.ValidateToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token")
	}

	return auth.SetClaims(ctx, claims), nil
}

func (g *GerenciaService) Migrate(ctx context.Context, req *pb.MigrateReq) (*pb.MigrateRes, error) {

	// TODO add claims/auth
	version := req.GetVersao()

	if err := migration.Migrate(ctx, version); err != nil {
		if errors.As(err, &migrate.ErrNoChange) {
			g.log.Infow("service", "migration", "no change in migration")
		} else {
			g.log.Errorw("migrate", "ERROR", err)
			return &pb.MigrateRes{}, err
		}
	}

	g.log.Infow(fmt.Sprintf("migrate to version %d", version))

	return &pb.MigrateRes{}, nil
}

func (g *GerenciaService) CreateCamera(ctx context.Context, req *pb.CreateCameraReq) (*pb.CreateCameraRes, error) {
	cam := camera.FromProto(req.Camera)

	// claims, err := auth.GetClaims(ctx)
	// if err != nil {
	// 	g.log.Errorw("auth", "claims missing from context", err)
	// 	return &pb.CreateCameraRes{}, errors.New("claims missing from context")
	// }

	camID, err := g.cameraStore.Create(ctx, cam)
	if err != nil {
		g.log.Errorw("create camera", "ERROR", err)
		return &pb.CreateCameraRes{}, fmt.Errorf("create: %w", err)
	}

	return &pb.CreateCameraRes{CameraId: camID}, nil
}

func (g *GerenciaService) ReadCamera(ctx context.Context, req *pb.ReadCameraReq) (*pb.ReadCameraRes, error) {

	cam, err := g.cameraStore.QueryByID(ctx, req.GetCameraId())
	if err != nil {
		g.log.Errorw("query camera", "ERROR", err)
		return &pb.ReadCameraRes{}, fmt.Errorf("query: %w", err)
	}

	return &pb.ReadCameraRes{Camera: cam.ToProto()}, err
}

func (g *GerenciaService) ReadCameras(ctx context.Context, req *pb.ReadCamerasReq) (*pb.ReadCamerasRes, error) {

	query := req.GetQuery()
	pageNumber := int(req.GetPageNumber())
	rowsPerPage := int(req.GetRowsPerPage())

	cameras, err := g.cameraStore.Query(ctx, query, pageNumber, rowsPerPage)
	if err != nil {
		g.log.Errorw("query cameras", "ERROR", err)
		return &pb.ReadCamerasRes{}, fmt.Errorf("query: %w", err)
	}

	return &pb.ReadCamerasRes{Cameras: cameras.ToProto()}, nil
}

func (g *GerenciaService) UpdateCamera(ctx context.Context, req *pb.UpdateCameraReq) (*pb.UpdateCameraRes, error) {

	cam := camera.FromProto(req.Camera)

	if err := g.cameraStore.Update(ctx, cam); err != nil {
		g.log.Errorw("update camera", "ERROR", err)
		return &pb.UpdateCameraRes{}, fmt.Errorf("update: %w", err)
	}
	return &pb.UpdateCameraRes{}, nil
}

func (g *GerenciaService) DeleteCamera(ctx context.Context, req *pb.DeleteCameraReq) (*pb.DeleteCameraRes, error) {

	for _, c := range req.GetCameraId() {
		if err := g.cameraStore.Delete(ctx, c); err != nil {
			g.log.Errorw("delete camera", "ERROR", err)
			return &pb.DeleteCameraRes{}, fmt.Errorf("delete: %w", err)
		}
	}

	return &pb.DeleteCameraRes{}, nil
}

func (g *GerenciaService) CreateUsuario(ctx context.Context, req *pb.CreateUsuarioReq) (*pb.CreateUsuarioRes, error) {
	usuario := usuario.FromProto(req.Usuario)

	// claims, err := auth.GetClaims(ctx)
	// if err != nil {
	// 	g.log.Errorw("auth", "claims missing from context", err)
	// 	return &pb.CreateCameraRes{}, errors.New("claims missing from context")
	// }

	usuarioID, err := g.usuarioStore.Create(ctx, usuario)
	if err != nil {
		g.log.Errorw("create usuario", "ERROR", err)
		return &pb.CreateUsuarioRes{}, fmt.Errorf("create: %w", err)
	}

	return &pb.CreateUsuarioRes{UsuarioId: usuarioID}, nil
}

func (g *GerenciaService) ReadUsuario(ctx context.Context, req *pb.ReadUsuarioReq) (*pb.ReadUsuarioRes, error) {

	usuario, err := g.usuarioStore.QueryByID(ctx, req.GetUsuarioId())
	if err != nil {
		g.log.Errorw("query usuario", "ERROR", err)
		return &pb.ReadUsuarioRes{}, fmt.Errorf("query: %w", err)
	}

	return &pb.ReadUsuarioRes{Usuario: usuario.ToProto()}, err
}

func (g *GerenciaService) ReadUsuarios(ctx context.Context, req *pb.ReadUsuariosReq) (*pb.ReadUsuariosRes, error) {

	query := req.GetQuery()
	pageNumber := int(req.GetPageNumber())
	rowsPerPage := int(req.GetRowsPerPage())

	usuarios, err := g.usuarioStore.Query(ctx, query, pageNumber, rowsPerPage)
	if err != nil {
		g.log.Errorw("query usuarios", "ERROR", err)
		return &pb.ReadUsuariosRes{}, fmt.Errorf("query: %w", err)
	}

	return &pb.ReadUsuariosRes{Usuarios: usuarios.ToProto()}, nil
}

func (g *GerenciaService) UpdateUsuario(ctx context.Context, req *pb.UpdateUsuarioReq) (*pb.UpdateUsuarioRes, error) {

	usuario := usuario.FromProto(req.Usuario)

	if err := g.usuarioStore.Update(ctx, usuario); err != nil {
		g.log.Errorw("update usuario", "ERROR", err)
		return &pb.UpdateUsuarioRes{}, fmt.Errorf("update: %w", err)
	}

	return &pb.UpdateUsuarioRes{}, nil
}

func (g *GerenciaService) DeleteUsuario(ctx context.Context, req *pb.DeleteUsuarioReq) (*pb.DeleteUsuarioRes, error) {

	for _, u := range req.GetUsuarioId() {
		if err := g.usuarioStore.Delete(ctx, u); err != nil {
			g.log.Errorw("delete usuario", "ERROR", err)
			return &pb.DeleteUsuarioRes{}, fmt.Errorf("delete: %w", err)
		}
	}

	return &pb.DeleteUsuarioRes{}, nil
}

func (g *GerenciaService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {

	fmt.Println("chegou aqui")

	claims, err := g.usuarioStore.Authenticate(ctx, req.GetEmail(), req.GetSenha())
	if err != nil {
		if errors.As(err, &database.ErrNotFound) {
			g.log.Errorw("usuario not found", "ERROR", err)
			return &pb.LoginRes{}, err
		} else if errors.As(err, &database.ErrAuthenticationFailure) {
			g.log.Errorw("authenticate usuario", "ERROR", err)
			return &pb.LoginRes{}, err
		}
		g.log.Errorw("authenticating", "ERROR", err)
		return &pb.LoginRes{}, err
	}

	var tkn struct {
		Token string `json:"token"`
	}
	tkn.Token, err = g.auth.GenerateToken(claims)
	if err != nil {
		return &pb.LoginRes{}, fmt.Errorf("generating token: %w", err)
	}

	return &pb.LoginRes{AccessToken: tkn.Token}, nil // TODO devolver o Token em contexto eu acho - no header
}
