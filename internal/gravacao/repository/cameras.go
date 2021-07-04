package repository

import (
	"context"

	"github.com/filipeandrade6/vigia-go/internal/gravacao/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CamerasRepository interface {
	Save(camera *models.Camera) error
	GetByID(id string) (camera *models.Camera, err error)
	GetByIP(id string) (camera *models.Camera, err error)
	GetAll() ([]*models.Camera, error)
	Update(camera *models.Camera) error
	Delete(id string) error
}

func GetByID(p *pgxpool.Pool, id string) (models.Camera, error) {
	query := `SELECT * FROM "camera" WHERE id = $1`

	var cam models.Camera

	if err := p.QueryRow(context.Background(), query, id).Scan(
		&cam.ID,
		&cam.IP,
		&cam.Descricao,
		&cam.Porta,
		&cam.Canal,
		&cam.UsuarioCamera,
		&cam.SenhaCamera,
		&cam.Cidade,
		&cam.Geolocalizacao,
		&cam.Marca,
		&cam.Modelo,
		&cam.Informacao,
	); err != nil {
		return models.Camera{}, err
	}

	return cam, nil
}
