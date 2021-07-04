package models

type Camera struct {
	ID             int
	IP             string
	Descricao      string
	Porta          int
	Canal          int
	UsuarioCamera  string
	SenhaCamera    string
	Cidade         string
	Geolocalizacao string
	Marca          string
	Modelo         string
	Informacao     string
}
