package server

import (
	"context"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/config"
	"github.com/filipeandrade6/vigia-go/internal/sys/database"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type gerenciaServer struct {
	pb.UnimplementedGerenciaServer
	cfg config.Gerencia
	db  config.DB
}

func NewGerenciaServer(config config) *grpc.Server {
	grpcGerenciaServer := grpc.NewServer()
	pb.RegisterGerenciaServer(grpcGerenciaServer, &gerenciaServer{cfg: config.Gerencia, db: config.DB})

	return grpcGerenciaServer
}

func (s *gerenciaServer) Migrate(ctx context.Context, in *pb.MigrateReq) (*pb.MigrateResp, error) {
	db, err := database.Open(s.config)

	return &pb.MigrateResp{}, nil
}

func (s *gerenciaServer) RegistrarServidorGravacao(ctx context.Context, req *pb.RegistrarServidorGravacaoReq) (*pb.RegistrarServidorGravacaoResp, error) {
	// sg := &models.ServidorGravacao{}
	// sg.FromProtobuf(req)
	// // TODO tratar a informação
	// sg.ID = "1"
	// sg.Status = "OK"
	// return sg.ToProtobuf(), nil
	return nil, nil
}

func (s *gerenciaServer) ConfigBancoDeDados(ctx context.Context, req *pb.ConfigBancoDeDadosReq) (*pb.ConfigBancoDeDadosResp, error) {
	return &pb.ConfigBancoDeDadosResp{
		Host:         viper.GetString("DB_HOST"),
		Port:         int32(viper.GetInt("DB_PORT")),
		User:         viper.GetString("DB_USER"),
		Password:     viper.GetString("DB_PASS"),
		Dbname:       viper.GetString("DB_NAME"),
		Poolmaxconns: int32(viper.GetInt("DB_POOL")),
	}, nil
}
