package models

// TODO duplicado com o gerencia

import pb "github.com/filipeandrade6/vigia-go/internal/api/v1"

type ServidorGravacao struct {
	ID         string
	EnderecoIP string
	Porta      int32
}

func (s *ServidorGravacao) ToProtobuf() *pb.RegistrarServidorGravacaoReq {
	return &pb.RegistrarServidorGravacaoReq{
		Id:         s.ID,
		EnderecoIp: s.EnderecoIP,
		Porta:      s.Porta,
	}
}

func (s *ServidorGravacao) FromProtobuf(sv *pb.RegistrarServidorGravacaoResp) {
	s.ID = sv.GetId()
}
