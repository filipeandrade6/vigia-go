package usuario

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/sys/auth"
	"github.com/filipeandrade6/vigia-go/internal/sys/database"
	"github.com/filipeandrade6/vigia-go/internal/sys/validate"
	"github.com/golang-jwt/jwt/v4"
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

func (s Store) Create(ctx context.Context, claims auth.Claims, usuario Usuario) (string, error) {
	if !claims.Authorized(auth.RoleAdmin) {
		return "", database.ErrForbidden
	}

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

func (s Store) Query(ctx context.Context, claims auth.Claims, query string, pageNumber int, rowsPerPage int) (Usuarios, error) {
	if !claims.Authorized(auth.RoleAdmin) {
		return Usuarios{}, database.ErrForbidden
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
		if errors.As(err, &database.ErrNotFound) { // TODO verificar se esse database.ErrNotFound funciona com ponteiro...
			return Usuarios{}, database.ErrNotFound
		}
		return Usuarios{}, fmt.Errorf("selecting usuarios: %w", err)
	}

	return usuarios, nil
}

func (s Store) QueryByID(ctx context.Context, claims auth.Claims, usuarioID string) (Usuario, error) {
	if err := validate.CheckID(usuarioID); err != nil {
		return Usuario{}, database.ErrInvalidID
	}

	if !claims.Authorized(auth.RoleAdmin) {
		return Usuario{}, database.ErrForbidden
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
		if errors.As(err, &database.ErrNotFound) {
			return Usuario{}, database.ErrNotFound
		}
		return Usuario{}, fmt.Errorf("selecting usuarioID[%q]: %w", usuarioID, err)
	}

	return usuario, nil
}

func (s Store) Update(ctx context.Context, claims auth.Claims, usuario Usuario) error {
	if err := validate.CheckID(usuario.UsuarioID); err != nil {
		return database.ErrInvalidID
	}

	if !claims.Authorized(auth.RoleAdmin) {
		return database.ErrForbidden
	}

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
		"funcao" = :funcao
	WHERE
		usuario_id = :usuario_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, usuario); err != nil {
		return fmt.Errorf("updating usuarioID[%s]: %w", usuario.UsuarioID, err)
	}

	return nil
}

func (s Store) Delete(ctx context.Context, claims auth.Claims, usuarioID string) error {
	if err := validate.CheckID(usuarioID); err != nil {
		return database.ErrInvalidID
	}

	if !claims.Authorized(auth.RoleAdmin) && claims.Subject != usuarioID {
		return database.ErrForbidden
	}

	data := struct {
		UsuarioID string `db:"usuario_id"`
	}{
		UsuarioID: usuarioID,
	}

	const q = `
	DELETE FROM
		usuarios
	WHERE
		usuario_id = :usuario_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("deleting usuarioID[%s]: %w", usuarioID, err)
	}

	return nil
}

func (s Store) Authenticate(ctx context.Context, email, senha string) (auth.Claims, error) {
	data := struct {
		Email string `db:"email"`
	}{
		Email: email,
	}

	fmt.Println("chegou aqui")

	const q = `
	SELECT
		*
	FROM
		usuarios
	WHERE
		email = :email`

	var usuario Usuario
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &usuario); err != nil {
		if errors.As(err, &database.ErrNotFound) {
			return auth.Claims{}, database.ErrNotFound
		}
		return auth.Claims{}, fmt.Errorf("selecting usuario[%q]: %w", email, err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(senha)); err != nil {
		return auth.Claims{}, database.ErrAuthenticationFailure
	}

	claims := auth.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "service project",
			Subject:   usuario.UsuarioID,
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
		},
		Roles: usuario.Funcao,
	}

	return claims, nil
}
