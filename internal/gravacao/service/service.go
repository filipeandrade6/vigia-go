package service

import (
	gerenciaGRPC "github.com/filipeandrade6/vigia-go/internal/gerencia/grpc"
	"go.uber.org/zap"
)

type GravacaoService struct {
	log            *zap.SugaredLogger
	gerenciaClient *gerenciaGRPC.GerenciaClient
}

func NewGravacaoService(log *zap.SugaredLogger, gerenciaClient *gerenciaGRPC.GerenciaClient) *GravacaoService {
	return &GravacaoService{
		log:            log,
		gerenciaClient: gerenciaClient,
	}
}
