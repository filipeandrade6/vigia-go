package processo

import (
	"unsafe"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/core/processo/db"
)

// TODO colcoar campos agregados e data de criacao e edicao

type Processo struct {
	ProcessoID         string `json:"processo_id"`
	ServidorGravacaoID string `json:"servidor_gravacao_id"`
	CameraID           string `json:"camera_id"`
	Processador        int    `json:"processador"`
	Adaptador          int    `json:"adaptador"`
	Execucao           bool   `json:"execucao"`
}

type NewProcesso struct {
	ServidorGravacaoID string `json:"servidor_gravacao_id" validate:"required"`
	CameraID           string `json:"camera_id" validate:"required"`
	Processador        int    `json:"processador" validate:"required"`
	Adaptador          int    `json:"adaptador" validate:"required"`
	Execucao           bool   `json:"execucao"`
}

type UpdateProcesso struct {
	ServidorGravacaoID *string `json:"servidor_gravacao_id" validate:"omitempty"`
	CameraID           *string `json:"camera_id" validate:"omitempty"`
	Processador        *int    `json:"processador" validate:"omitempty"`
	Adaptador          *int    `json:"adaptador" validate:"omitempty"`
	Execucao           *bool   `json:"execucao" validate:"omitempty"`
}

// =============================================================================

func toProcesso(dbPrc db.Processo) Processo {
	p := (*Processo)(unsafe.Pointer(&dbPrc))
	return *p
}

func toProcessoSlice(dbPrcs []db.Processo) []Processo {
	prcs := make([]Processo, len(dbPrcs))
	for i, dbPrc := range dbPrcs {
		prcs[i] = toProcesso(dbPrc)
	}
	return prcs
}

// =============================================================================

func (p Processo) ToProto() *pb.Processo {
	return &pb.Processo{
		ProcessoId:         p.ProcessoID,
		ServidorGravacaoId: p.ServidorGravacaoID,
		CameraId:           p.CameraID,
		Processador:        int32(p.Processador),
		Adaptador:          int32(p.Adaptador),
		Execucao:           p.Execucao,
	}
}

func FromProto(p *pb.Processo) Processo {
	return Processo{
		ProcessoID:         p.GetProcessoId(),
		ServidorGravacaoID: p.GetServidorGravacaoId(),
		CameraID:           p.GetCameraId(),
		Processador:        int(p.GetProcessador()),
		Adaptador:          int(p.GetAdaptador()),
		Execucao:           p.GetExecucao(),
	}
}

type Processos []Processo

func (p Processos) ToProto() []*pb.Processo {
	var prcs []*pb.Processo

	for _, prc := range p {
		prcs = append(prcs, prc.ToProto())
	}

	return prcs
}
