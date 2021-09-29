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

type ServidoresGravacao []ServidorGravacao

func (s ServidoresGravacao) ToProto() []*pb.ServidorGravacao {
	var svs []*pb.ServidorGravacao

	for _, sv := range s {
		svs = append(svs, sv.ToProto())
	}

	return svs
}

func ServidoresGravacaoFromProto(s []*pb.ServidorGravacao) ServidoresGravacao {
	var svs ServidoresGravacao

	for _, sv := range s {
		svs = append(svs, FromProto(sv))
	}

	return svs
}
