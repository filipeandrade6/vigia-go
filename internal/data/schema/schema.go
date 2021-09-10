// // TODO alterar isso aqui...
// // Package schema contains the database schema, migrations and seeding data.
package schema

import (
	"context"
	// _ "embed" // Calls init function.
	"embed"
	"fmt"

	// TODO no lugar do darwin utilizar o go-migrations
	"github.com/ardanlabs/service/business/sys/database"
	"github.com/jmoiron/sqlx"

	"github.com/golang-migrate/migrate/v4"
	"github.com/johejo/golang-migrate-extra/source/iofs"
)

// var (
// 	//go:embed sql/schema.sql
// 	schemaDoc string

// 	//go:embed sql/seed.sql
// 	seedDoc string

// 	//go:embed sql/delete.sql
// 	deleteDoc string
// )

//go:embed migrations/*.sql
var fs embed.FS

// Migrate attempts to bring the schema for db up to date with the migrations
// defined in this package.
func Migrate(ctx context.Context, db *sqlx.DB) error {
	if err := database.StatusCheck(ctx, db); err != nil {
		return fmt.Errorf("status check database: %w", err)
	}

	d, err := iofs.New(fs, "migrations")
	if err != nil {
		return err // TODO arrumar isso aqui
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, "postgres://postgres@localhost/postgres?sslmode=disable")
	if err != nil {
		return fmt.Errorf("construct postgres driver: %w", err)
	}

	err = m.Up()
	if err != nil {
		return fmt.Errorf("uping") // TODO arrumar isso aqui
	}

	return nil

	// driver, err := darwin.NewGenericDriver(db.DB, darwin.PostgresDialect{})
	// if err != nil {
	// 	return fmt.Errorf("construct darwin driver: %w", err)
	// }

	// d := darwin.New(driver, darwin.ParseMigrations(schemaDoc))
	// return d.Migrate()
}

// Seed runs the set of seed-data queries against db. The queries are ran in a
// transaction and rolled back if any fail.
// func Seed(ctx context.Context, db *sqlx.DB) error {
// 	if err := database.StatusCheck(ctx, db); err != nil {
// 		return fmt.Errorf("status check database: %w", err)
// 	}

// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	if _, err := tx.Exec(seedDoc); err != nil {
// 		if err := tx.Rollback(); err != nil {
// 			return err
// 		}
// 		return err
// 	}

// 	return tx.Commit()
// }

// // DeleteAll runs the set of Drop-table queries against db. The queries are ran in a
// // transaction and rolled back if any fail.
// func DeleteAll(db *sqlx.DB) error {
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	if _, err := tx.Exec(deleteDoc); err != nil {
// 		if err := tx.Rollback(); err != nil {
// 			return err
// 		}
// 		return err
// 	}

// 	return tx.Commit()
// }
