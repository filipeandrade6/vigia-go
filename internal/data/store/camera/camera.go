package camera

import (
	"context"
	"fmt"

	"github.com/filipeandrade6/vigia-go/internal/sys/auth"
	"github.com/filipeandrade6/vigia-go/internal/sys/database"
	"github.com/filipeandrade6/vigia-go/internal/sys/validate"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Store struct {
	log *zap.SugaredLogger
	db  *sqlx.DB
}

func NewStore(log *zap.SugaredLogger, db *sqlx.DB) Store {
	return Store{
		log: log,
		db:  db,
	}
}

func (s Store) Create(ctx context.Context, claims auth.Claims, cam Camera) (string, error) {

	if !claims.Authorized(auth.RoleAdmin, auth.RoleManager) {
		return "", database.ErrForbidden
	}

	c := Camera{
		CameraID:       validate.GenerateID(),
		Descricao:      cam.Descricao,
		EnderecoIP:     cam.EnderecoIP,
		Porta:          cam.Porta,
		Canal:          cam.Canal,
		Usuario:        cam.Usuario,
		Senha:          cam.Senha,
		Geolocalizacao: cam.Geolocalizacao,
	}

	const q = `
	INSERT INTO cameras
		(camera_id, descricao, endereco_ip, porta, canal, usuario, senha, geolocalizacao)
	VALUES
		(:camera_id, :descricao, :endereco_ip, :porta, :canal, :usuario, :senha, :geolocalizacao)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, c); err != nil {
		return "", fmt.Errorf("inserting camera: %w", err)
	}

	return c.CameraID, nil
}

func (s Store) Query(ctx context.Context, query string, pageNumber int, rowsPerPage int) (Cameras, error) {
	data := struct {
		Query       string `db:"query"`
		Offset      int    `db:"offset"`
		RowsPerPage int    `db:"rows_per_page"`
	}{
		Query:       query,
		Offset:      (pageNumber - 1) * rowsPerPage,
		RowsPerPage: rowsPerPage,
	}

	// TODO verificar sqlinjection
	const q = `
	SELECT
		*
	FROM
		cameras
	WHERE
		(camera_id, descricao, endereco_ip, porta, canal, usuario, senha, geolocalizacao)
	ILIKE
		(%:query%)
	ORDER BY
		descricao
	OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY`

	var cams Cameras
	if err := database.NamedQuerySlice(ctx, s.log, s.db, q, data, &cams); err != nil {
		if err == database.ErrNotFound {
			return nil, database.ErrNotFound
		}
		return nil, fmt.Errorf("selecting cameras: %w", err)
	}

	return cams, nil
}

func (s Store) QueryByID(ctx context.Context, cameraID string) (Camera, error) {
	if err := validate.CheckID(cameraID); err != nil {
		return Camera{}, database.ErrInvalidID
	}

	data := struct {
		CameraID string `db:"camera_id"`
	}{
		CameraID: cameraID,
	}

	const q = `
	SELECT
		*
	FROM
		cameras
	WHERE
		camera_id = :camera_id`

	var cam Camera
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &cam); err != nil {
		if err == database.ErrNotFound {
			return Camera{}, database.ErrNotFound
		}
		return Camera{}, fmt.Errorf("selecting camera ID[%q]: %w", data.CameraID, err)
	}

	return cam, nil
}

func (s Store) Update(ctx context.Context, cam Camera) error {
	if err := validate.CheckID(cam.CameraID); err != nil {
		return database.ErrInvalidID
	}

	// TODO implementar validate.Check
	// if err := validate.Check(camera); err != nil {
	// 	return fmt.Errorf("validating data: %w", err)
	// }

	// c, err := s.QueryByID(ctx, cam.CameraID)
	// if err != nil {
	// 	return fmt.Errorf("updating cameraID[%s]: %w", cam.CameraID, err)
	// }

	const q = `
	UPDATE
		cameras
	SET
		"descricao" = :descricao,
		"endereco_ip" = :endereco_ip,
		"porta" = :porta,
		"canal" = :canal,
		"usuario" = :usuario,
		"senha" = :senha,
		"geolocalizacao" = :geolocalizacao
	WHERE
		camera_id = :camera_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, cam); err != nil {
		return fmt.Errorf("updating cameraID[%s]: %w", cam.CameraID, err)
	}

	return nil
}

func (s Store) Delete(ctx context.Context, cameraID string) error {
	if err := validate.CheckID(cameraID); err != nil {
		return database.ErrInvalidID
	}

	data := struct {
		CameraID string `db:"camera_id"`
	}{
		CameraID: cameraID,
	}

	const q = `
	DELETE FROM
		cameras
	WHERE
		camera_id = :camera_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("deleting cameraID[%s]: %w", cameraID, err)
	}

	return nil
}
