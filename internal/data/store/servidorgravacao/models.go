package servidorgravacao

import (
	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
)

type ServidorGravacao struct {
	ServidorGravacaoID string `db:"servidor_gravacao_id"`
	EnderecoIP         string `db:"endereco_ip"`
	Porta              int    `db:"porta"`
	Host               string `db:"host"`
}

func (s ServidorGravacao) ToProto() *pb.ServidorGravacao {
	return &pb.ServidorGravacao{
		ServidorGravacaoId: s.ServidorGravacaoID,
		EnderecoIp:         s.EnderecoIP,
		Porta:              int32(s.Porta),
		Host:               s.Host,
	}
}

func FromProto(s *pb.ServidorGravacao) ServidorGravacao {
	return ServidorGravacao{
		ServidorGravacaoID: s.GetServidorGravacaoId(),
		EnderecoIP:         s.GetEnderecoIp(),
		Porta:              int(s.GetPorta()),
		Host:               s.GetHost(),
	}
}
