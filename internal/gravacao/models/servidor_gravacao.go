package models

import pb "github.com/filipeandrade6/vigia-go/internal/api/v1"

type ServidorGravacao struct {
	ID       int32
	Endereço string
	Porta    int32
}

func (s *ServidorGravacao) ToProtobuf() *pb.GravacaoConfigReq {
	return &pb.GravacaoConfigReq{
		ServidorGravacao:      s.Endereço,
		PortaServidorGravacao: s.Porta,
	}
}

// func (s *ServidorGravacao) FromProtobuf(servidor *pb.ServidorGravacao) *pb.GravacaoConfigResp {
// 	s.ID = servidor.GetId()
// }
