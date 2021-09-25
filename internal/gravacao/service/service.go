package service

import (
	"github.com/filipeandrade6/vigia-go/internal/grpc/gerencia"
	"go.uber.org/zap"
)

type GravacaoService struct {
	log            *zap.SugaredLogger
	gerenciaClient *gerencia.GerenciaClient
}

func NewGravacaoService(log *zap.SugaredLogger, gerenciaClient *gerencia.GerenciaClient) *GravacaoService {
	return &GravacaoService{
		log:            log,
		gerenciaClient: gerenciaClient,
	}
}
