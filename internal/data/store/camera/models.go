package camera

import "time"

type Camera struct {
	ID              int
	Nome            string
	IP              string
	Porta           int
	Canal           int
	Usuario         string
	Senha           string
	Regiao          string
	Geolocalizacao  string
	Marca           string
	Modelo          string
	Informacao      string
	Consumo         string // Aggregate field showing total of storatge consuption
	DataCriacao     time.Time
	DataAtualizacao time.Time
}
