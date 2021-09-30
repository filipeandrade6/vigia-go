package camera

import (
	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
)

type Camera struct {
	CameraID       string `db:"camera_id"`
	Descricao      string `db:"descricao"`
	EnderecoIP     string `db:"endereco_ip" validate:""`
	Porta          int    `db:"porta" validate:""`
	Canal          int    `db:"canal" validate:""`
	Usuario        string `db:"usuario" validate:""`
	Senha          string `db:"senha" validate:""`
	Geolocalizacao string `db:"geolocalizacao" validate:""`
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
	}
}

type Cameras []Camera

func (c Cameras) ToProto() []*pb.Camera {
	var cams []*pb.Camera

	for _, cam := range c {
		cams = append(cams, cam.ToProto())
	}

	return cams
}

func CamerasFromProto(c []*pb.Camera) Cameras {
	var cams Cameras

	for _, cam := range c {
		cams = append(cams, FromProto(cam))
	}

	return cams
}
