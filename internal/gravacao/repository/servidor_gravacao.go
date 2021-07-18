package repository

import (
	"context"

	"github.com/filipeandrade6/vigia-go/internal/gravacao/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ServidorGravacaoRepository interface {
	SaveServidorGravacao(sv *models.ServidorGravacao) (*models.ServidorGravacao, error)
	GetServidorGravacaoByID(id int32) (*models.ServidorGravacao, error)
	GetServidorGravacaoByEndereco(endereco string) (*models.ServidorGravacao, error)
	GetAllServidorGravacao() ([]*models.ServidorGravacao, error)
	UpdateServidorGravacao(sv *models.ServidorGravacao) error
	DeleteServidorGravacao(id int32) error
}

func SaveServidorGravacao(p *pgxpool.Pool, sv *models.ServidorGravacao) (*models.ServidorGravacao, error) {
	query := `INSERT INTO servidor_gravacao (endereco, porta) VALUES ($1, $2) RETURNING id`

	var id int32

	if err := p.QueryRow(context.Background(), query, sv).Scan(&id); err != nil {
		return &models.ServidorGravacao{}, err
	}

	return &models.ServidorGravacao{ID: id}, nil
}

func GetServidorGravacaoByID(p *pgxpool.Pool, id int32) (*models.ServidorGravacao, error) {
	query := `SELECT * FROM servidor_gravacao WHERE id = $1`

	var sv *models.ServidorGravacao

	if err := p.QueryRow(context.Background(), query, id).Scan(
		&sv.ID,
		&sv.Endereço,
		&sv.Porta,
	); err != nil {
		return &models.ServidorGravacao{}, err
	}

	return sv, nil
}

func GetServidorGravacaoByEndereco(p *pgxpool.Pool, endereco string) (*models.ServidorGravacao, error) {
	query := `SELECT * FROM "servidor_gravacao" WHERE endereco = $1`

	var sv *models.ServidorGravacao

	if err := p.QueryRow(context.Background(), query, endereco).Scan(
		&sv.ID,
		&sv.Endereço,
		&sv.Porta,
	); err != nil {
		return &models.ServidorGravacao{}, err
	}

	return sv, nil
}

// func GetAllServidorGravacao(p *pgxpool.Pool) ([]*models.ServidorGravacao, error) {
// 	query := `SELECT * F`
// 	var svs []*models.ServidorGravacao
// 	if err := p.Query()

// 	return nil, nil
// }
