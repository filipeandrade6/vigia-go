package servidorgravacao

import (
	"time"

	pb "github.com/filipeandrade6/vigia-go/internal/api"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ServidorGravacao struct {
	ServidorGravacaoID string    `db:"servidor_gravacao_id"`
	EnderecoIP         string    `db:"endereco_ip"`
	Porta              int       `db:"porta"`
	Host               string    `db:"host"`
	CriadoEm           time.Time `db:"criado_em"`
	EditadoEm          time.Time `db:"editado_em"`
}

func (s ServidorGravacao) ToProto() *pb.ServidorGravacao {
	return &pb.ServidorGravacao{
		ServidorGravacaoId: s.ServidorGravacaoID,
		EnderecoIp:         s.EnderecoIP,
		Porta:              int32(s.Porta),
		Host:               s.Host,
		CriadoEm:           timestamppb.New(s.CriadoEm),
		EditadoEm:          timestamppb.New(s.EditadoEm),
	}
}

func FromProto(s *pb.ServidorGravacao) ServidorGravacao {
	return ServidorGravacao{
		ServidorGravacaoID: s.GetServidorGravacaoId(),
		EnderecoIP:         s.GetEnderecoIp(),
		Porta:              int(s.GetPorta()),
		Host:               s.GetHost(),
		CriadoEm:           s.GetCriadoEm().AsTime(),
		EditadoEm:          s.GetEditadoEm().AsTime(),
	}
}
