package migration

import (
	"context"
	"embed"
	"fmt"

	"github.com/spf13/viper"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/johejo/golang-migrate-extra/source/iofs"
)

//go:embed sql/*.sql
var fs embed.FS

// Migrate attempts to bring the schema for db up to date with the migrations
// defined in this package.

// TODO: verificar se esta conectado o banco de dados - Status Check e seed
func Migrate(ctx context.Context) error {
	d, err := iofs.New(fs, "sql")
	if err != nil {
		return err // TODO arrumar isso aqui
	}

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		viper.GetString("VIGIA_DB_USER"),
		viper.GetString("VIGIA_DB_PASSWORD"),
		viper.GetString("VIGIA_DB_HOST"),
		viper.GetString("VIGIA_DB_NAME"),
		viper.GetString("VIGIA_DB_SSLMODE"),
	)

	m, err := migrate.NewWithSourceInstance("iofs", d, dbURL)
	if err != nil {
		return fmt.Errorf("construct postgres driver: %w", err)
	}

	err = m.Up()
	if err != nil {
		return err
	}

	return nil
}
