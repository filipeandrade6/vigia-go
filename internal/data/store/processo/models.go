package processo

import (
	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
)

type Processo struct {
	ProcessoID         string `db:"processo_Id"`
	ServidorGravacaoID string `db:"servidor_gravacao_id"`
	CameraID           string `db:"camera_id"`
	Processador        int    `db:"processador"`
	Adaptador          int    `db:"adaptador"`
	Execucao           bool   `db:"execucao"`
}

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

func ProcessosFromProto(p []*pb.Processo) Processos {
	var prcs Processos

	for _, prc := range p {
		prcs = append(prcs, FromProto(prc))
	}

	return prcs
}
