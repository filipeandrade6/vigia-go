package servidorgravacao

import (
	"context"
	"errors"
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

func (s *Store) Create(ctx context.Context, claims auth.Claims, sv ServidorGravacao) (string, error) {
	if !claims.Authorized(auth.RoleAdmin) {
		return "", database.ErrForbidden
	}

	sv.ServidorGravacaoID = validate.GenerateID()

	const q = `
	INSERT INTO servidores_gravacao
		(servidor_gravacao_id, endereco_ip, porta, host)
	VALUES
		(:servidor_gravacao_id, :endereco_ip, :porta, :host)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, sv); err != nil {
		return "", fmt.Errorf("inserting servidor gravacao: %w", err)
	}

	return sv.ServidorGravacaoID, nil
}

func (s Store) QueryByID(ctx context.Context, claims auth.Claims, svID string) (ServidorGravacao, error) {
	if err := validate.CheckID(svID); err != nil {
		return ServidorGravacao{}, database.ErrInvalidID
	}

	if !claims.Authorized(auth.RoleAdmin) {
		return ServidorGravacao{}, database.ErrForbidden
	}

	data := struct {
		CameraID string `db:"servidor_gravacao_id"`
	}{
		CameraID: svID,
	}

	const q = `
	SELECT
		*
	FROM
		servidores_gravacao
	WHERE
		servidor_gravacao_id = :servidor_gravacao_id`

	var sv ServidorGravacao
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &sv); err != nil {
		if errors.As(err, &database.ErrNotFound) {
			return ServidorGravacao{}, database.ErrNotFound
		}
		return ServidorGravacao{}, fmt.Errorf("selecting svID[%q]: %w", svID, err)
	}

	return sv, nil
}
