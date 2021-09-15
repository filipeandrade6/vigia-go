package service

import (
	"github.com/filipeandrade6/vigia-go/internal/gerencia/grpc/client"
	"go.uber.org/zap"
)

type GravacaoService struct {
	log            *zap.SugaredLogger
	gerenciaClient *client.GerenciaClient
}

func NewGravacaoService(log *zap.SugaredLogger, gerenciaClient *client.GerenciaClient) *GravacaoService {
	return &GravacaoService{
		log:            log,
		gerenciaClient: gerenciaClient,
	}
}
