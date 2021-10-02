package camera

import (
	"unsafe"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/core/camera/db"
)

// TODO colcoar campos agregados e data de criacao e edicao

type Camera struct {
	CameraID   string
	Descricao  string
	EnderecoIP string
	Porta      int
	Canal      int
	Usuario    string
	Senha      string
	Latitude   string
	Longitude  string
}

type NewCamera struct {
	Descricao  string `validate:"required"`
	EnderecoIP string `validate:"required,ip"`
	Porta      int    `validate:"required,gte=1,lte=65536"`
	Canal      int    `validate:"required,gte=0,lte=10"`
	Usuario    string `validate:"required"`
	Senha      string `validate:"required"`
	Latitude   string `validate:"required,latitude"`
	Longitude  string `validate:"required,longitude"`
}

type UpdateCamera struct {
	Descricao  *string `validate:"omitempty"`
	EnderecoIP *string `validate:"omitempty,ip"`
	Porta      *int    `validate:"omitempty,gte=1,lte=65536"`
	Canal      *int    `validate:"omitempty,gte=0,lte=10"`
	Usuario    *string `validate:"omitempty"`
	Senha      *string `validate:"omitempty"`
	Latitude   *string `validate:"omitempty,latitude"`
	Longitude  *string `validate:"omitempty,longitude"`
}

// =============================================================================

func toCamera(dbCam db.Camera) Camera {
	c := (*Camera)(unsafe.Pointer(&dbCam))
	return *c
}

func toCameraSlice(dbCams []db.Camera) []Camera {
	cams := make([]Camera, len(dbCams))
	for i, dbCam := range dbCams {
		cams[i] = toCamera(dbCam)
	}
	return cams
}

// =============================================================================

func (c Camera) ToProto() *pb.Camera {
	return &pb.Camera{
		CameraId:   c.CameraID,
		Descricao:  c.Descricao,
		EnderecoIp: c.EnderecoIP,
		Porta:      int32(c.Porta),
		Canal:      int32(c.Canal),
		Usuario:    c.Usuario,
		Senha:      c.Senha,
		Latitude:   c.Latitude,
		Longitude:  c.Longitude,
	}
}

func FromProto(c *pb.Camera) Camera {
	return Camera{
		CameraID:   c.GetCameraId(),
		Descricao:  c.GetDescricao(),
		EnderecoIP: c.GetEnderecoIp(),
		Porta:      int(c.GetPorta()),
		Canal:      int(c.GetCanal()),
		Usuario:    c.GetUsuario(),
		Senha:      c.GetSenha(),
		Latitude:   c.GetLatitude(),
		Longitude:  c.GetLongitude(),
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
