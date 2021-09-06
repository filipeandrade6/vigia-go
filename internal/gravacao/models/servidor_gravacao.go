package models

// // TODO duplicado com o gerencia

// import (
// 	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
// 	"github.com/jackc/pgx/v4/pgxpool"
// )

// // https://www.alexedwards.net/blog/organising-database-access

// type ServidorGravacao struct {
// 	ID         string
// 	EnderecoIP string
// 	Porta      int32
// }

// type ServidorGravacaoModel struct {
// 	DB *pgxpool.Pool
// }

// func (m ServidorGravacaoModel) SaveServidorGravacao(sv ServidorGravacao) ([]ServidorGravacao, error) {

// 	rows, err := m.DB.Query()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var servidoresGravacao []ServidorGravacao

// 	for rows.Next() {
// 		var sv ServidorGravacao

// 		if err := rows.Scan(&sv.ID, &sv.EnderecoIP, &sv.Porta); err != nil {
// 			return nil, err
// 		}

// 		servidoresGravacao = append(servidoresGravacao, sv)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return servidoresGravacao, nil
// }

// func (s *ServidorGravacao) ToProtobuf() *pb.RegistrarServidorGravacaoReq {
// 	return &pb.RegistrarServidorGravacaoReq{
// 		Id:         s.ID,
// 		EnderecoIp: s.EnderecoIP,
// 		Porta:      s.Porta,
// 	}
// }

// func (s *ServidorGravacao) FromProtobuf(sv *pb.RegistrarServidorGravacaoResp) {
// 	s.ID = sv.GetId()
// }
