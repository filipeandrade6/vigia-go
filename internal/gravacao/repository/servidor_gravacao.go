package repository

import (
	"context"

	"github.com/filipeandrade6/vigia-go/internal/gravacao/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

// TODO criar estrutura que tem o pgxpool como campo
// TODO criar uma interface para query https://github.com/johanbrandhorst/grpc-postgres/tree/master/users

type ServidorGravacaoRepository interface {
	SaveServidorGravacao(sv *models.ServidorGravacao) (*models.ServidorGravacao, error)
	GetAllServidorGravacao() ([]*models.ServidorGravacao, error)
	GetServidorGravacaoByID(id string) (*models.ServidorGravacao, error)
	// GetServidorGravacaoByEndereco(endereco string) (*models.ServidorGravacao, error)
	// UpdateServidorGravacao(sv *models.ServidorGravacao) error
	DeleteServidorGravacao(id string) error
}

func SaveServidorGravacao(p *pgxpool.Pool, sv *models.ServidorGravacao) (*models.ServidorGravacao, error) {
	query := `INSERT INTO servidores_gravacao (id, endereco_ip, porta) VALUES ($1, $2, $3) RETURNING id`

	var id string

	if err := p.QueryRow(context.Background(), query, sv).Scan(&id); err != nil {
		return &models.ServidorGravacao{}, err
	}

	return &models.ServidorGravacao{ID: id}, nil
}

func GetServidorGravacaoByID(p *pgxpool.Pool, id string) (*models.ServidorGravacao, error) {
	query := `SELECT * FROM servidores_gravacao WHERE id = $1`

	var sv *models.ServidorGravacao

	if err := p.QueryRow(context.Background(), query, id).Scan(
		&sv.ID,
		&sv.EnderecoIP,
		&sv.Porta,
	); err != nil {
		return &models.ServidorGravacao{}, err
	}

	return sv, nil
}

func GetAllServidorGravacao(p *pgxpool.Pool) ([]*models.ServidorGravacao, error) {
	query := `SELECT * FROM servidores_gravacao`

	var svs []*models.ServidorGravacao

	rows, err := p.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sv *models.ServidorGravacao
		err := rows.Scan(&sv.ID, &sv.EnderecoIP, &sv.Porta)
		if err != nil {
			return nil, err // TODO arrumar isso aqui
		}
		svs = append(svs, sv)
	}
	if err := rows.Err(); err != nil {
		return nil, err // TODO arrumar isso aqui
	}

	return svs, nil
}

func DeleteServidorGravacao(p *pgxpool.Pool, id string) error {
	query := `DELETE FROM servidores_gravacao WHERE id = $1 RETURNING id`

	if err := p.QueryRow(context.Background(), query, id).Scan(); err != nil {
		return err // TODO tratar ErrNoRows
	}
	return nil
}

// func GetServidorGravacaoByEndereco(p *pgxpool.Pool, endereco string) (*models.ServidorGravacao, error) {
// 	query := `SELECT * FROM "servidor_gravacao" WHERE endereco = $1`

// 	var sv *models.ServidorGravacao

// 	if err := p.QueryRow(context.Background(), query, endereco).Scan(
// 		&sv.ID,
// 		&sv.Endereço,
// 		&sv.Porta,
// 	); err != nil {
// 		return &models.ServidorGravacao{}, err
// 	}

// 	return sv, nil
// }
