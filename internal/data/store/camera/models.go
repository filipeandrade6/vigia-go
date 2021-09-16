package camera

import "time"

// TODO aggregate fields - consumo, leituras, tempo de atividade

type Camera struct {
	CameraID       string    `db:"camera_id"`
	Descricao      string    `db:"descricao"`
	EnderecoIP     string    `db:"endereco_ip"`
	Porta          int       `db:"porta"`
	Canal          int       `db:"canal"`
	Usuario        string    `db:"usuario"`
	Senha          string    `db:"senha"`
	Geolocalizacao string    `db:"geolocalizao"`
	CriadoEm       time.Time `db:"criado_em"`
	EditadoEm      time.Time `db:"editado_em"`
}
