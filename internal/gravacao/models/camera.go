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

// * pb é o pacote do protobuffer gerado a partir dos arquivos em messages
// func (c *Camera) ToProtobuffer() *pb.Camera{
// 	return &pb.Camera{
// 		Id: c.ID,
// 		Name: c.Name,
// 		Email: c.Email,
// 		...
// 	}
// }

// func (c *Camera) FromProtobuffer(camera *pb.Camera) {
// 	u.Id = camera.GetId()
// 	u.Name = camera.GetName()
// 	...
// }
