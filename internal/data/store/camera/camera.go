package camera

import (
	"context"
	"fmt"
	"time"

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

func (s Store) Create(ctx context.Context, cam Camera, now time.Time) (Camera, error) {
	c := Camera{
		CameraID:       validate.GenerateID(),
		Descricao:      cam.Descricao,
		EnderecoIP:     cam.EnderecoIP,
		Porta:          cam.Porta,
		Canal:          cam.Canal,
		Usuario:        cam.Usuario,
		Senha:          cam.Senha,
		Geolocalizacao: cam.Geolocalizacao,
		CriadoEm:       now,
	}

	const q = `
	INSERT INTO cameras
		(camera_id, descricao, endereco_ip, porta, canal, usuario, senha, geolocalizacao, criado_em)
	VALUES
		(:camera_id, :descricao, :endereco_ip, :porta, :canal, :usuario, :senha, :geolocalizacao, :criado_em)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, c); err != nil {
		return Camera{}, fmt.Errorf("inserting camera: %w", err)
	}

	return cam, nil
}

func (s Store) Query(ctx context.Context, pageNumber int, rowsPerPage int) ([]Camera, error) {
	data := struct {
		Offset      int `db:"offset"`
		RowsPerPage int `db:"rows_per_page"`
	}{
		Offset:      (pageNumber - 1) * rowsPerPage,
		RowsPerPage: rowsPerPage,
	}

	const q = `
	SELECT
		*
	FROM
		cameras
	ORDER BY
		camera_id
	OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY`

	var cams []Camera
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

// func (s Store) Update(ctx context.Context, cameraID string, cam Camera, now time.Time) error {
// 	if err := validate.CheckID(cameraID); err != nil {
// 		return database.ErrInvalidID
// 	}
// 	if err := validate.Check(camera); err != nil {
// 		return fmt.Errorf("validating data: %w", err)
// 	}

// 	c, err := s.QueryByID(ctx, cameraID)
// 	if err != nil {
// 		return fmt.Errorf("updating camera cameraID[%s]: %w", cameraID, err)
// 	}

// 	c.EditadoEm = now

// 	const q = `
// 	UPDATE
// 		cameras
// 	SET
// 		"descricao" = :descricao,
// 		"ip" = :ip,
// 		"porta" = :porta,
// 		"canal" = :canal,
// 		"usuario" = :usuario,
// 		"senha" = :senha,
// 		"geolocalizacao" = :geolocalizacao,
// 		"editado_em" = :editado_em,
// 	WHERE
// 		camera_id = :camera_id`

// 	if err := database.NamedExecContext(ctx, s.log, s.db, q, c); err != nil {
// 		return fmt.Errorf("updating cameraID[%s]: %w", cameraID, err)
// 	}

// 	return nil
// }

func (s Store) Delete(ctx context.Context, cameraID string) error {
	if err := validate.CheckID(cameraID); err != nil {
		return database.ErrInvalidID
	}

	// TODO entender abaixo
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
