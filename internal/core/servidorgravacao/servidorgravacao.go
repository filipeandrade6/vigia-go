package servidorgravacao

import (
	"context"
	"errors"
	"fmt"

	"github.com/filipeandrade6/vigia-go/internal/core/servidorgravacao/db"
	"github.com/filipeandrade6/vigia-go/internal/sys/auth"
	"github.com/filipeandrade6/vigia-go/internal/sys/database"
	"github.com/filipeandrade6/vigia-go/internal/sys/validate"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var (
	ErrNotFound  = errors.New("servidor de gravacao not found")
	ErrInvalidID = errors.New("ID is not in its proper from")
)

type Core struct {
	store db.Store
}

func NewCore(log *zap.SugaredLogger, sqlxDB *sqlx.DB) Core {
	return Core{
		store: db.NewStore(log, sqlxDB),
	}
}

func (c Core) Create(ctx context.Context, nsv NewServidorGravacao) (ServidorGravacao, error) {
	if err := validate.Check(nsv); err != nil {
		return ServidorGravacao{}, fmt.Errorf("validating data: %w", err)
	}

	dbSV := db.ServidorGravacao{
		ServidorGravacaoID: validate.GenerateID(),
		EnderecoIP:         nsv.EnderecoIP,
		Porta:              nsv.Porta,
	}

	if err := c.store.Create(ctx, dbSV); err != nil {
		return ServidorGravacao{}, fmt.Errorf("create: %w", err)
	}

	return toServidorGravacao(dbSV), nil
}

func (c Core) Query(ctx context.Context, claims auth.Claims, query string, pageNumber int, rowsPerPage int) (ServidoresGravacao, error) {
	if !claims.Authorized(auth.RoleAdmin) {
		return ServidoresGravacao{}, database.ErrForbidden
	}

	data := struct {
		Query       string `db:"query"`
		Offset      int    `db:"offset"`
		RowsPerPage int    `db:"rows_per_page"`
	}{
		Query:       fmt.Sprintf("%%%s%%", query),
		Offset:      (pageNumber - 1) * rowsPerPage,
		RowsPerPage: rowsPerPage,
	}

	const q = `
	SELECT
		*
	FROM
		servidores_gravacao
	WHERE
		CONCAT(servidor_gravacao_id, endereco_ip, porta, host)
	ILIKE
		:query
	ORDER BY
		host
	OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY`

	var svs ServidoresGravacao
	if err := database.NamedQuerySlice(ctx, s.log, s.db, q, data, &svs); err != nil {
		if errors.As(err, &database.ErrNotFound) {
			return ServidoresGravacao{}, database.ErrNotFound
		}
		return ServidoresGravacao{}, fmt.Errorf("selecting servidores de gravacao: %w", err)
	}

	return svs, nil
}

func (c Core) QueryByID(ctx context.Context, claims auth.Claims, svID string) (ServidorGravacao, error) {
	if err := validate.CheckID(svID); err != nil {
		return ServidorGravacao{}, database.ErrInvalidID
	}

	if !claims.Authorized(auth.RoleAdmin) {
		return ServidorGravacao{}, database.ErrForbidden
	}

	data := struct {
		SvID string `db:"servidor_gravacao_id"`
	}{
		SvID: svID,
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

func (c Core) Update(ctx context.Context, claims auth.Claims, sv ServidorGravacao) error {
	if !claims.Authorized(auth.RoleAdmin) {
		return database.ErrForbidden
	}

	if err := validate.CheckID(sv.ServidorGravacaoID); err != nil {
		return database.ErrInvalidID
	}

	const q = `
	UPDATE
		servidores_gravacao
	SET
		"endereco_ip" = :endereco_ip,
		"porta" = :porta,
		"host" = :host
	WHERE
		servidor_gravacao_id = :servidor_gravacao_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, sv); err != nil {
		return fmt.Errorf("updating svID[%s]: %w", sv.ServidorGravacaoID, err)
	}

	return nil
}

func (c Core) Delete(ctx context.Context, claims auth.Claims, svID string) error {

	// TODO verificar se o servidor não está em execução, perguntar, etc...

	if !claims.Authorized(auth.RoleAdmin) {
		return database.ErrForbidden
	}

	if err := validate.CheckID(svID); err != nil {
		return database.ErrInvalidID
	}

	data := struct {
		SvID string `db:"servidor_gravacao_id"`
	}{
		SvID: svID,
	}

	const q = `
	DELETE FROM
		servidores_gravacao
	WHERE
		servidor_gravacao_id = :servidor_gravacao_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("deleting svID[%s]: %w", svID, err)
	}

	return nil
}
