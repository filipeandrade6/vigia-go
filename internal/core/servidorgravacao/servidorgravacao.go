package servidorgravacao

import (
	"context"
	"errors"
	"fmt"

	"github.com/filipeandrade6/vigia-go/internal/core/servidorgravacao/db"
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

func (c Core) Update(ctx context.Context, svID string, up UpdateServidorGravacao) error {
	if err := validate.CheckID(svID); err != nil {
		return ErrInvalidID
	}

	dbSV, err := c.store.QueryByID(ctx, svID)
	if err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			return ErrNotFound
		}
		return fmt.Errorf("updating servidor de gravacao svID[%s]: %w", svID, err)
	}

	if up.EnderecoIP != nil {
		dbSV.EnderecoIP = *up.EnderecoIP
	}
	if up.Porta != nil {
		dbSV.Porta = *up.Porta
	}

	if err := c.store.Update(ctx, dbSV); err != nil {
		return fmt.Errorf("update: %w", err)
	}

	return nil
}

func (c Core) Delete(ctx context.Context, svID string) error {
	if err := validate.CheckID(svID); err != nil {
		return ErrInvalidID
	}

	if err := c.store.Delete(ctx, svID); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

func (c Core) Query(ctx context.Context, query string, pageNumber int, rowsPerPage int) (ServidoresGravacao, error) {
	dbSVs, err := c.store.Query(ctx, query, pageNumber, rowsPerPage)
	if err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("query: %w", err)
	}

	return toServidorGravacaoSlice(dbSVs), nil
}

func (c Core) QueryByID(ctx context.Context, svID string) (ServidorGravacao, error) {
	if err := validate.CheckID(svID); err != nil {
		return ServidorGravacao{}, ErrInvalidID
	}

	dbSV, err := c.store.QueryByID(ctx, svID)
	if err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			return ServidorGravacao{}, ErrNotFound
		}
		return ServidorGravacao{}, fmt.Errorf("query: %w", err)
	}

	return toServidorGravacao(dbSV), nil
}
