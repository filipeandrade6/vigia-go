package camera

import (
	"time"

	pb "github.com/filipeandrade6/vigia-go/internal/api"
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

func ToProto(camera Camera) *pb.CreateCameraReq {

}

func FromProto(camera *pb.CreateCameraReq) Camera {

}
