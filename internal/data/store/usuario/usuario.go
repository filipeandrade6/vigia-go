package usuario

import (
	"context"
	"fmt"

	"github.com/filipeandrade6/vigia-go/internal/sys/database"
	"github.com/filipeandrade6/vigia-go/internal/sys/validate"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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

func (s Store) Create(ctx context.Context, usuario Usuario) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("generating password hash: %w", err)
	}

	usuario.UsuarioID = validate.GenerateID()
	usuario.Senha = string(hash)

	const q = `
	INSERT INTO usuarios
		(usuario_id, email, senha_hash, funcao)
	VALUES
		(:usuario_id, :email, :senha_hash, :funcao)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, usuario); err != nil {
		return "", fmt.Errorf("inserting usuario: %w", err)
	}

	return usuario.UsuarioID, nil
}

func (s Store) Query(ctx context.Context, query string, pageNumber int, rowsPerPage int) (Usuarios, error) {
	data := struct {
		Query       string `db:"query"`
		Offset      int    `db:"offset"`
		RowsPerPage int    `db:"rows_per_page"`
	}{
		Query:       fmt.Sprintf("%%%s%%", query),
		Offset:      (pageNumber - 1) * rowsPerPage,
		RowsPerPage: rowsPerPage,
	}

	// TODO verificar sqlinjection
	const q = `
	SELECT
		*
	FROM
		usuarios
	WHERE
		CONCAT(usuario_id, email, senha_hash, funcao)
	ILIKE
		:query
	ORDER BY
		email
	OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY`

	var usuarios Usuarios
	if err := database.NamedQuerySlice(ctx, s.log, s.db, q, data, &usuarios); err != nil {
		if err == database.ErrNotFound {
			return Usuarios{}, database.ErrNotFound
		}
		return Usuarios{}, fmt.Errorf("selecting usuarios: %w", err)
	}

	return usuarios, nil
}

func (s Store) QueryByID(ctx context.Context, usuarioID string) (Usuario, error) {
	if err := validate.CheckID(usuarioID); err != nil {
		return Usuario{}, database.ErrInvalidID
	}

	data := struct {
		UsuarioID string `db:"usuario_id"`
	}{
		UsuarioID: usuarioID,
	}

	const q = `
	SELECT
		*
	FROM
		usuarios
	WHERE
		usuario_id = :usuario_id`

	var usuario Usuario
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &usuario); err != nil {
		if err == database.ErrNotFound {
			return Usuario{}, database.ErrNotFound
		}
		return Usuario{}, fmt.Errorf("selecting usuarioID[%q]: %w", usuarioID, err)
	}

	return usuario, nil
}

func (s Store) Update(ctx context.Context, usuario Usuario) error {
	if err := validate.CheckID(usuario.UsuarioID); err != nil {
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

	hash, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("generating password hash: %w", err)
	}

	usuario.Senha = string(hash)

	const q = `
	UPDATE
		usuarios
	SET
		"email" = :email,
		"senha_hash" = :senha_hash,
		"funcao" = :funcao,
		"canal" = :canal,
	WHERE
		usuario_id = :usuario_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, usuario); err != nil {
		return fmt.Errorf("updating usuarioID[%s]: %w", usuario.UsuarioID, err)
	}

	return nil
}

func (s Store) Delete(ctx context.Context, usuarioID string) error {
	if err := validate.CheckID(usuarioID); err != nil {
		return database.ErrInvalidID
	}

	data := struct {
		UsuarioID string `db:"usuario_id"`
	}{
		UsuarioID: usuarioID,
	}

	const q = `
	DELETE FROM
		usuario
	WHERE
		usuario_id = :usuario_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("deleting usuarioID[%s]: %w", usuarioID, err)
	}

	return nil
}
