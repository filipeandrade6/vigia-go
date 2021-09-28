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
	"github.com/filipeandrade6/vigia-go/internal/sys/validate"

	"github.com/golang-migrate/migrate/v4"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO registrar no log os erros - ver como fazer corretamente...

// TODO criar campo servicos - e registrar os servicos para não colocar tudo em um?

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
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("authentication", "ERROR", err)
		}
	}()

	// TODO se o servico chamado for Login - so encaminha o contexto
	fmt.Println(fullMethodName)
	return ctx, nil

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	fmt.Println(token)

	claims, err := g.auth.ValidateToken(token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid auth token")
	}

	return auth.SetClaims(ctx, claims), nil
}

// =========================================================================
// Migrate

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

// =========================================================================
// Usuario

func (g *GerenciaService) CreateUsuario(ctx context.Context, req *pb.CreateUsuarioReq) (*pb.CreateUsuarioRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("create usuario", "ERROR", err)
		}
	}()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return &pb.CreateUsuarioRes{}, status.Error(codes.Unauthenticated, "claims missing from context")
	}

	usuario := usuario.FromProto(req.Usuario)

	usuarioID, err := g.usuarioStore.Create(ctx, claims, usuario)
	if err != nil {
		if validate.Cause(err) == database.ErrForbidden {
			return &pb.CreateUsuarioRes{}, status.Error(codes.PermissionDenied, err.Error())
		}
		return &pb.CreateUsuarioRes{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateUsuarioRes{UsuarioId: usuarioID}, nil
}

func (g *GerenciaService) ReadUsuario(ctx context.Context, req *pb.ReadUsuarioReq) (*pb.ReadUsuarioRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("read usuario", "ERROR", err)
		}
	}()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return &pb.ReadUsuarioRes{}, status.Error(codes.Unauthenticated, "claims missing from context")
	}

	usuario, err := g.usuarioStore.QueryByID(ctx, claims, req.GetUsuarioId())
	if err != nil {
		switch validate.Cause(err) {
		case database.ErrInvalidID:
			return &pb.ReadUsuarioRes{}, status.Error(codes.InvalidArgument, err.Error())
		case database.ErrForbidden:
			return &pb.ReadUsuarioRes{}, status.Error(codes.PermissionDenied, err.Error())
		case database.ErrNotFound:
			return &pb.ReadUsuarioRes{}, status.Error(codes.NotFound, err.Error())
		default:
			return &pb.ReadUsuarioRes{}, status.Error(codes.Internal, err.Error())
		}
	}

	return &pb.ReadUsuarioRes{Usuario: usuario.ToProto()}, err
}

func (g *GerenciaService) ReadUsuarios(ctx context.Context, req *pb.ReadUsuariosReq) (*pb.ReadUsuariosRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("read usuarios", "ERROR", err)
		}
	}()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return &pb.ReadUsuariosRes{}, status.Error(codes.Unauthenticated, "claims missing from context")
	}

	query := req.GetQuery()
	pageNumber := int(req.GetPageNumber())
	rowsPerPage := int(req.GetRowsPerPage())

	usuarios, err := g.usuarioStore.Query(ctx, claims, query, pageNumber, rowsPerPage)
	if err != nil {
		switch validate.Cause(err) {
		case database.ErrForbidden:
			return &pb.ReadUsuariosRes{}, status.Error(codes.PermissionDenied, err.Error())
		case database.ErrNotFound:
			return &pb.ReadUsuariosRes{}, status.Error(codes.NotFound, err.Error())
		default:
			return &pb.ReadUsuariosRes{}, status.Error(codes.Internal, err.Error())
		}
	}
	return &pb.ReadUsuariosRes{Usuarios: usuarios.ToProto()}, nil
}

func (g *GerenciaService) UpdateUsuario(ctx context.Context, req *pb.UpdateUsuarioReq) (*pb.UpdateUsuarioRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("update usuario", "ERROR", err)
		}
	}()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return &pb.UpdateUsuarioRes{}, status.Error(codes.Unauthenticated, "claims missing from context")
	}

	usuario := usuario.FromProto(req.Usuario)

	if err := g.usuarioStore.Update(ctx, claims, usuario); err != nil {
		switch validate.Cause(err) {
		case database.ErrInvalidID:
			return &pb.UpdateUsuarioRes{}, status.Error(codes.InvalidArgument, err.Error())
		case database.ErrForbidden:
			return &pb.UpdateUsuarioRes{}, status.Error(codes.PermissionDenied, err.Error())
		case database.ErrNotFound:
			return &pb.UpdateUsuarioRes{}, status.Error(codes.NotFound, err.Error())
		default:
			return &pb.UpdateUsuarioRes{}, status.Error(codes.Internal, err.Error())
		}
	}

	return &pb.UpdateUsuarioRes{}, nil
}

func (g *GerenciaService) DeleteUsuario(ctx context.Context, req *pb.DeleteUsuarioReq) (*pb.DeleteUsuarioRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("delete usuario", "ERROR", err)
		}
	}()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return &pb.DeleteUsuarioRes{}, status.Error(codes.Unauthenticated, "claims missing from context")
	}

	for _, u := range req.GetUsuarioId() {
		if err := g.usuarioStore.Delete(ctx, claims, u); err != nil {
			switch validate.Cause(err) {
			case database.ErrInvalidID:
				return &pb.DeleteUsuarioRes{}, status.Error(codes.InvalidArgument, err.Error())
			case database.ErrForbidden:
				return &pb.DeleteUsuarioRes{}, status.Error(codes.PermissionDenied, err.Error())
			default:
				return &pb.DeleteUsuarioRes{}, status.Error(codes.Internal, err.Error())
			}
		}
	}

	return &pb.DeleteUsuarioRes{}, nil
}

func (g *GerenciaService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("login usuario", "ERROR", err)
		}
	}()

	claims, err := g.usuarioStore.Authenticate(ctx, req.GetEmail(), req.GetSenha())
	if err != nil {
		switch validate.Cause(err) {
		case database.ErrNotFound:
			return &pb.LoginRes{}, status.Error(codes.NotFound, err.Error())
		case database.ErrAuthenticationFailure:
			return &pb.LoginRes{}, status.Error(codes.Unauthenticated, err.Error())
		default:
			return &pb.LoginRes{}, status.Error(codes.Internal, err.Error())
		}
	}

	var tkn struct {
		Token string `json:"token"`
	}
	tkn.Token, err = g.auth.GenerateToken(claims)
	if err != nil {
		return &pb.LoginRes{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.LoginRes{AccessToken: tkn.Token}, nil // TODO devolver o Token em contexto eu acho - no header
}

// =========================================================================
// Camera

func (g *GerenciaService) CreateCamera(ctx context.Context, req *pb.CreateCameraReq) (*pb.CreateCameraRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("create camera", "ERROR", err)
		}
	}()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return &pb.CreateCameraRes{}, status.Error(codes.Unauthenticated, "claims missing from context")
	}

	cam := camera.FromProto(req.Camera)

	camID, err := g.cameraStore.Create(ctx, claims, cam)
	if err != nil {
		if validate.Cause(err) == database.ErrForbidden {
			return &pb.CreateCameraRes{}, status.Error(codes.PermissionDenied, err.Error())
		}
		return &pb.CreateCameraRes{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateCameraRes{CameraId: camID}, nil
}

func (g *GerenciaService) ReadCamera(ctx context.Context, req *pb.ReadCameraReq) (*pb.ReadCameraRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("read camera", "ERROR", err)
		}
	}()

	cam, err := g.cameraStore.QueryByID(ctx, req.GetCameraId())
	if err != nil {
		switch validate.Cause(err) {
		case database.ErrInvalidID:
			return &pb.ReadCameraRes{}, status.Error(codes.InvalidArgument, err.Error())
		case database.ErrNotFound:
			return &pb.ReadCameraRes{}, status.Error(codes.NotFound, err.Error())
		default:
			return &pb.ReadCameraRes{}, status.Error(codes.Internal, err.Error())
		}
	}

	return &pb.ReadCameraRes{Camera: cam.ToProto()}, nil
}

func (g *GerenciaService) ReadCameras(ctx context.Context, req *pb.ReadCamerasReq) (*pb.ReadCamerasRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("read cameras", "ERROR", err)
		}
	}()

	query := req.GetQuery()
	pageNumber := int(req.GetPageNumber())
	rowsPerPage := int(req.GetRowsPerPage())

	cameras, err := g.cameraStore.Query(ctx, query, pageNumber, rowsPerPage)
	if err != nil {
		if validate.Cause(err) == database.ErrNotFound {
			return &pb.ReadCamerasRes{}, status.Error(codes.NotFound, err.Error())
		}
		return &pb.ReadCamerasRes{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.ReadCamerasRes{Cameras: cameras.ToProto()}, nil
}

func (g *GerenciaService) UpdateCamera(ctx context.Context, req *pb.UpdateCameraReq) (*pb.UpdateCameraRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("update camera", "ERROR", err)
		}
	}()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return &pb.UpdateCameraRes{}, status.Error(codes.Unauthenticated, "claims missing from context")
	}

	cam := camera.FromProto(req.Camera)

	if err := g.cameraStore.Update(ctx, claims, cam); err != nil {
		switch validate.Cause(err) {
		case database.ErrForbidden:
			return &pb.UpdateCameraRes{}, status.Error(codes.PermissionDenied, err.Error())
		case database.ErrInvalidID:
			return &pb.UpdateCameraRes{}, status.Error(codes.InvalidArgument, err.Error())
		default:
			return &pb.UpdateCameraRes{}, status.Error(codes.Internal, err.Error())
		}
	}
	return &pb.UpdateCameraRes{}, nil
}

func (g *GerenciaService) DeleteCamera(ctx context.Context, req *pb.DeleteCameraReq) (*pb.DeleteCameraRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("delete camera", "ERROR", err)
		}
	}()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return &pb.DeleteCameraRes{}, status.Error(codes.Unauthenticated, "claims missing from context")
	}

	for _, c := range req.GetCameraId() {
		if err := g.cameraStore.Delete(ctx, claims, c); err != nil {
			switch validate.Cause(err) {
			case database.ErrForbidden:
				return &pb.DeleteCameraRes{}, status.Error(codes.PermissionDenied, err.Error())
			case database.ErrInvalidID:
				return &pb.DeleteCameraRes{}, status.Error(codes.InvalidArgument, err.Error())
			default:
				return &pb.DeleteCameraRes{}, status.Error(codes.Internal, err.Error())
			}
		}
	}

	return &pb.DeleteCameraRes{}, nil
}
