package db

type ServidorGravacao struct {
	ServidorGravacaoID string `db:"servidor_gravacao_id"`
	EnderecoIP         string `db:"endereco_ip"`
	Porta              int    `db:"porta"`
	Host               string `db:"host"`
}
