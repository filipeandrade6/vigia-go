package service

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/database/migration"

	"github.com/filipeandrade6/vigia-go/internal/core/camera"
	"github.com/filipeandrade6/vigia-go/internal/core/servidorgravacao"
	"github.com/filipeandrade6/vigia-go/internal/core/usuario"
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

type FrontendService struct {
	pb.UnimplementedFrontendServer
	log                  *zap.SugaredLogger
	auth                 *auth.Auth
	dbCfg                database.Config
	cameraCore           camera.Core
	usuarioCore          usuario.Core
	servidorGravacaoCore servidorgravacao.Core
}

func NewFrontendService(
	log *zap.SugaredLogger,
	auth *auth.Auth,
	dbCfg database.Config,
	cameraCore camera.Core,
	usuarioCore usuario.Core,
	servidorGravacaoCore servidorgravacao.Core,
) *FrontendService {
	return &FrontendService{
		log:                  log,
		auth:                 auth,
		dbCfg:                dbCfg,
		cameraCore:           cameraCore,
		usuarioCore:          usuarioCore,
		servidorGravacaoCore: servidorGravacaoCore,
	}
}

func (g *FrontendService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("authentication", "ERROR", err)
		}
	}()

	// TODO se o servico chamado for Login - so encaminha o contexto
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

func (g *FrontendService) Migrate(ctx context.Context, req *pb.MigrateReq) (*pb.MigrateRes, error) {

	// TODO add claims/auth
	version := req.GetVersao()

	dbURL := "" // TODO arrumar

	if err := migration.Migrate(ctx, version, dbURL); err != nil {
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

func (g *FrontendService) CreateUsuario(ctx context.Context, req *pb.CreateUsuarioReq) (*pb.CreateUsuarioRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("create usuario", "ERROR", err)
		}
	}()

	usr := usuario.FromProto(req.Usuario)
	nu := usuario.NewUsuario{
		Email:  usr.Email,
		Funcao: usr.Funcao,
		Senha:  usr.Senha,
	}

	savedUsr, err := g.usuarioCore.Create(ctx, nu)
	if err != nil {
		if validate.Cause(err) == database.ErrForbidden {
			return &pb.CreateUsuarioRes{}, status.Error(codes.PermissionDenied, err.Error())
		}
		return &pb.CreateUsuarioRes{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateUsuarioRes{UsuarioId: savedUsr.UsuarioID}, nil
}

func (g *FrontendService) ReadUsuario(ctx context.Context, req *pb.ReadUsuarioReq) (*pb.ReadUsuarioRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("read usuario", "ERROR", err)
		}
	}()

	usuario, err := g.usuarioCore.QueryByID(ctx, req.GetUsuarioId())
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

func (g *FrontendService) ReadUsuarios(ctx context.Context, req *pb.ReadUsuariosReq) (*pb.ReadUsuariosRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("read usuarios", "ERROR", err)
		}
	}()

	query := req.GetQuery()
	pageNumber := int(req.GetPageNumber())
	rowsPerPage := int(req.GetRowsPerPage())

	usuarios, err := g.usuarioCore.Query(ctx, query, pageNumber, rowsPerPage)
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

func (g *FrontendService) UpdateUsuario(ctx context.Context, req *pb.UpdateUsuarioReq) (*pb.UpdateUsuarioRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("update usuario", "ERROR", err)
		}
	}()

	upd := usuario.UpdateUsuario{
		Email:  req.GetEmail(),
		Funcao: req.GetFuncao(),
		Senha:  req.GetSenha(),
	}

	if err := g.usuarioCore.Update(ctx, upd); err != nil {
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

func (g *FrontendService) DeleteUsuario(ctx context.Context, req *pb.DeleteUsuarioReq) (*pb.DeleteUsuarioRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("delete usuario", "ERROR", err)
		}
	}()

	for _, u := range req.GetUsuarioId() {
		if err := g.usuarioCore.Delete(ctx, u); err != nil {
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

// func (g *FrontendService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
// 	var err error
// 	defer func() {
// 		if err != nil {
// 			g.log.Errorw("login usuario", "ERROR", err)
// 		}
// 	}()

// 	// claims, err := g.usuarioCore.Authenticate(ctx, req.GetEmail(), req.GetSenha())
// 	// if err != nil {
// 	// 	switch validate.Cause(err) {
// 	// 	case database.ErrNotFound:
// 	// 		return &pb.LoginRes{}, status.Error(codes.NotFound, err.Error())
// 	// 	case database.ErrAuthenticationFailure:
// 	// 		return &pb.LoginRes{}, status.Error(codes.Unauthenticated, err.Error())
// 	// 	default:
// 	// 		return &pb.LoginRes{}, status.Error(codes.Internal, err.Error())
// 	// 	}
// 	// }

// 	var tkn struct {
// 		Token string `json:"token"`
// 	}
// 	tkn.Token, err = g.auth.GenerateToken(claims)
// 	if err != nil {
// 		return &pb.LoginRes{}, status.Error(codes.Internal, err.Error())
// 	}

// 	return &pb.LoginRes{AccessToken: tkn.Token}, nil // TODO devolver o Token em contexto eu acho - no header
// }

// =========================================================================
// Camera

func (g *FrontendService) CreateCamera(ctx context.Context, req *pb.CreateCameraReq) (*pb.CreateCameraRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("create camera", "ERROR", err)
		}
	}()

	cam := camera.FromProto(req.Camera)
	nc := camera.NewCamera{
		Descricao:  cam.Descricao,
		EnderecoIP: cam.EnderecoIP,
		Porta:      cam.Porta,
		Canal:      cam.Canal,
		Usuario:    cam.Usuario,
		Senha:      cam.Senha,
		Latitude:   cam.Latitude,
		Longitude:  cam.Longitude,
	}

	savedCam, err := g.cameraCore.Create(ctx, nc)
	if err != nil {
		if validate.Cause(err) == database.ErrForbidden {
			return &pb.CreateCameraRes{}, status.Error(codes.PermissionDenied, err.Error())
		}
		return &pb.CreateCameraRes{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateCameraRes{CameraId: savedCam.CameraID}, nil
}

func (g *FrontendService) ReadCamera(ctx context.Context, req *pb.ReadCameraReq) (*pb.ReadCameraRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("read camera", "ERROR", err)
		}
	}()

	cam, err := g.cameraCore.QueryByID(ctx, req.GetCameraId())
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

func (g *FrontendService) ReadCameras(ctx context.Context, req *pb.ReadCamerasReq) (*pb.ReadCamerasRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("read cameras", "ERROR", err)
		}
	}()

	query := req.GetQuery()
	pageNumber := int(req.GetPageNumber())
	rowsPerPage := int(req.GetRowsPerPage())

	cameras, err := g.cameraCore.Query(ctx, query, pageNumber, rowsPerPage)
	if err != nil {
		if validate.Cause(err) == database.ErrNotFound {
			return &pb.ReadCamerasRes{}, status.Error(codes.NotFound, err.Error())
		}
		return &pb.ReadCamerasRes{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.ReadCamerasRes{Cameras: cameras.ToProto()}, nil
}

// func (g *FrontendService) UpdateCamera(ctx context.Context, req *pb.UpdateCameraReq) (*pb.UpdateCameraRes, error) {
// 	var err error
// 	defer func() {
// 		if err != nil {
// 			g.log.Errorw("update camera", "ERROR", err)
// 		}
// 	}()

// 	cam := camera.FromProto(req.Camera)
// 	upd := camera.UpdateCamera{
// 		Descricao:  &cam.Descricao,
// 		EnderecoIP: &cam.EnderecoIP,
// 		Porta:      &cam.Porta,
// 		Canal:      &cam.Canal,
// 		Usuario:    &cam.Usuario,
// 		Senha:      &cam.Senha,
// 		Latitude:   &cam.Latitude,
// 		Longitude:  &cam.Longitude,
// 	}

// 	if err := g.cameraCore.Update(ctx, cam.CameraID, upd); err != nil {
// 		switch validate.Cause(err) {
// 		case database.ErrForbidden:
// 			return &pb.UpdateCameraRes{}, status.Error(codes.PermissionDenied, err.Error())
// 		case database.ErrInvalidID:
// 			return &pb.UpdateCameraRes{}, status.Error(codes.InvalidArgument, err.Error())
// 		default:
// 			return &pb.UpdateCameraRes{}, status.Error(codes.Internal, err.Error())
// 		}
// 	}
// 	return &pb.UpdateCameraRes{}, nil
// }

func (g *FrontendService) DeleteCamera(ctx context.Context, req *pb.DeleteCameraReq) (*pb.DeleteCameraRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("delete camera", "ERROR", err)
		}
	}()

	for _, c := range req.GetCameraId() {
		if err := g.cameraCore.Delete(ctx, c); err != nil {
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

// =========================================================================
// Servidores Gravacao

func (g *FrontendService) CreateServidorGravacao(ctx context.Context, req *pb.CreateServidorGravacaoReq) (*pb.CreateServidorGravacaoRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("create servidor de gravacao", "ERROR", err)
		}
	}()

	sv := servidorgravacao.FromProto(req.ServidorGravacao)
	nsv := servidorgravacao.NewServidorGravacao{
		EnderecoIP: sv.EnderecoIP,
		Porta:      sv.Porta,
	}

	savedSV, err := g.servidorGravacaoCore.Create(ctx, nsv)
	if err != nil {
		if validate.Cause(err) == database.ErrForbidden {
			return &pb.CreateServidorGravacaoRes{}, status.Error(codes.PermissionDenied, err.Error())
		}
		return &pb.CreateServidorGravacaoRes{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateServidorGravacaoRes{ServidorGravacaoId: savedSV.ServidorGravacaoID}, nil
}

func (g *FrontendService) ReadServidorGravacao(ctx context.Context, req *pb.ReadServidorGravacaoReq) (*pb.ReadServidorGravacaoRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("read servidor de gravacao", "ERROR", err)
		}
	}()

	sv, err := g.servidorGravacaoCore.QueryByID(ctx, req.GetServidorGravacaoId())
	if err != nil {
		switch validate.Cause(err) {
		case database.ErrInvalidID:
			return &pb.ReadServidorGravacaoRes{}, status.Error(codes.InvalidArgument, err.Error())
		case database.ErrNotFound:
			return &pb.ReadServidorGravacaoRes{}, status.Error(codes.NotFound, err.Error())
		default:
			return &pb.ReadServidorGravacaoRes{}, status.Error(codes.Internal, err.Error())
		}
	}

	return &pb.ReadServidorGravacaoRes{ServidorGravacao: sv.ToProto()}, nil
}

func (g *FrontendService) ReadServidoresGravacao(ctx context.Context, req *pb.ReadServidoresGravacaoReq) (*pb.ReadServidoresGravacaoRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("read servidores de gravacao", "ERROR", err)
		}
	}()

	query := req.GetQuery()
	pageNumber := int(req.GetPageNumber())
	rowsPerPage := int(req.GetRowsPerPage())

	svs, err := g.servidorGravacaoCore.Query(ctx, query, pageNumber, rowsPerPage)
	if err != nil {
		if validate.Cause(err) == database.ErrNotFound {
			return &pb.ReadServidoresGravacaoRes{}, status.Error(codes.NotFound, err.Error())
		}
		return &pb.ReadServidoresGravacaoRes{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.ReadServidoresGravacaoRes{ServiceGravacao: svs.ToProto()}, nil
}

// func (g *FrontendService) UpdateServidorGravacao(ctx context.Context, req *pb.UpdateServidorGravacaoReq) (*pb.UpdateServidorGravacaoRes, error) {
// 	var err error
// 	defer func() {
// 		if err != nil {
// 			g.log.Errorw("update servidor de gravacao", "ERROR", err)
// 		}
// 	}()

// 	sv := servidorgravacao.FromProto(req.ServidorGravacao)
// 	upd := servidorgravacao.UpdateServidorGravacao{
// 		EnderecoIP: &sv.EnderecoIP,
// 		Porta:      &sv.Porta,
// 	}

// 	if err := g.servidorGravacaoCore.Update(ctx, sv.ServidorGravacaoID, upd); err != nil {
// 		switch validate.Cause(err) {
// 		case database.ErrForbidden:
// 			return &pb.UpdateServidorGravacaoRes{}, status.Error(codes.PermissionDenied, err.Error())
// 		case database.ErrInvalidID:
// 			return &pb.UpdateServidorGravacaoRes{}, status.Error(codes.InvalidArgument, err.Error())
// 		default:
// 			return &pb.UpdateServidorGravacaoRes{}, status.Error(codes.Internal, err.Error())
// 		}
// 	}
// 	return &pb.UpdateServidorGravacaoRes{}, nil
// }

func (g *FrontendService) DeleteServidorGravacao(ctx context.Context, req *pb.DeleteServidorGravacaoReq) (*pb.DeleteServidorGravacaoRes, error) {
	var err error
	defer func() {
		if err != nil {
			g.log.Errorw("delete servidor de gravacao", "ERROR", err)
		}
	}()

	for _, c := range req.GetServidorGravacaoId() {
		if err := g.servidorGravacaoCore.Delete(ctx, c); err != nil {
			switch validate.Cause(err) {
			case database.ErrForbidden:
				return &pb.DeleteServidorGravacaoRes{}, status.Error(codes.PermissionDenied, err.Error())
			case database.ErrInvalidID:
				return &pb.DeleteServidorGravacaoRes{}, status.Error(codes.InvalidArgument, err.Error())
			default:
				return &pb.DeleteServidorGravacaoRes{}, status.Error(codes.Internal, err.Error())
			}
		}
	}

	return &pb.DeleteServidorGravacaoRes{}, nil
}
