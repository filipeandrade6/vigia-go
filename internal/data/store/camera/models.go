package camera

import (
	"time"

	pb "github.com/filipeandrade6/vigia-go/internal/api"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// TODO aggregate fields - consumo, leituras, tempo de atividade

type Camera struct {
	CameraID       string    `db:"camera_id"`
	Descricao      string    `db:"descricao"`
	EnderecoIP     string    `db:"endereco_ip"`
	Porta          int       `db:"porta"`
	Canal          int       `db:"canal"`
	Usuario        string    `db:"usuario"`
	Senha          string    `db:"senha"`
	Geolocalizacao string    `db:"geolocalizacao"`
	CriadoEm       time.Time `db:"criado_em"`
	EditadoEm      time.Time `db:"editado_em"`
}

func (c Camera) ToProto() *pb.Camera {
	return &pb.Camera{
		CameraId:       c.CameraID,
		Descricao:      c.Descricao,
		EnderecoIp:     c.EnderecoIP,
		Porta:          int32(c.Porta),
		Canal:          int32(c.Canal),
		Usuario:        c.Usuario,
		Senha:          c.Senha,
		Geolocalizacao: c.Geolocalizacao,
		CriadoEm:       timestamppb.New(c.CriadoEm),
		EditadoEm:      timestamppb.New(c.EditadoEm),
	}
}

func FromProto(c *pb.Camera) Camera {
	return Camera{
		CameraID:       c.GetCameraId(),
		Descricao:      c.GetDescricao(),
		EnderecoIP:     c.GetEnderecoIp(),
		Porta:          int(c.GetPorta()),
		Canal:          int(c.GetCanal()),
		Usuario:        c.GetUsuario(),
		Senha:          c.GetSenha(),
		Geolocalizacao: c.GetGeolocalizacao(),
		CriadoEm:       c.GetCriadoEm().AsTime(),
		EditadoEm:      c.GetEditadoEm().AsTime(),
	}
}
