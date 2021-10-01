package servidorgravacao

import (
	"unsafe"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/core/servidorgravacao/db"
)

type ServidorGravacao struct {
	ServidorGravacaoID string
	EnderecoIP         string
	Porta              int
}

type NewServidorGravacao struct {
	EnderecoIP string `validate:"required,ip"`
	Porta      int    `validate:"required,gte=1,lte=65536"`
}

type UpdateServidorGravacao struct {
	EnderecoIP *string `validate:"omitempty,ip"`
	Porta      *int    `validate:"omitempty,gte=1,lte=65536"`
}

// =============================================================================

func toServidorGravacao(dbSV db.ServidorGravacao) ServidorGravacao {
	s := (*ServidorGravacao)(unsafe.Pointer(&dbSV))
	return *s
}

func toServidorGravacaoSlice(dbSVs []db.ServidorGravacao) []ServidorGravacao {
	svs := make([]ServidorGravacao, len(dbSVs))
	for i, dbSV := range dbSVs {
		svs[i] = toServidorGravacao(dbSV)
	}
	return svs
}

// =============================================================================

func (s ServidorGravacao) ToProto() *pb.ServidorGravacao {
	return &pb.ServidorGravacao{
		ServidorGravacaoId: s.ServidorGravacaoID,
		EnderecoIp:         s.EnderecoIP,
		Porta:              int32(s.Porta),
	}
}

func FromProto(s *pb.ServidorGravacao) ServidorGravacao {
	return ServidorGravacao{
		ServidorGravacaoID: s.GetServidorGravacaoId(),
		EnderecoIP:         s.GetEnderecoIp(),
		Porta:              int(s.GetPorta()),
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
