package processo

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

func (s Store) Create(ctx context.Context, claims auth.Claims, prc Processo) (string, error) {
	if !claims.Authorized(auth.RoleAdmin, auth.RoleManager) {
		return "", database.ErrForbidden
	}

	prc.ProcessoID = validate.GenerateID()

	const q = `
	INSERT INTO processos
		(processo_id, servidor_gravacao_id, camera_id, processador, adaptador, execucao)
	VALUES
		(:processo_id, :servidor_gravacao_id, :camera_id, :processador, :adaptador, :execucao)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, prc); err != nil {
		return "", fmt.Errorf("inserting processo: %w", err)
	}

	return prc.ProcessoID, nil
}
