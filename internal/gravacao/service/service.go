package service

import (
	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"go.uber.org/zap"
)

type GravacaoService struct {
	pb.UnimplementedGravacaoServer
	log *zap.SugaredLogger
}

func NewGravacaoService(log *zap.SugaredLogger) *GravacaoService {
	return &GravacaoService{
		log: log,
	}
}
