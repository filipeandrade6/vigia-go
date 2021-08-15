package database

// import (
// 	"database/sql"
// 	"embed"

// 	migrate "github.com/golang-migrate/migrate/v4"
// 	"github.com/golang-migrate/migrate/v4/database/postgres"
// )

// // TODO removido suporte para iofs pois estava quebrando o go mod tidy
// // https://github.com/golang-migrate/migrate/releases/tag/v4.14.1

// //go:embed migrations/*.sql
// var fs embed.FS

// // version defines the current migration version. This ensures the app
// // is always compatible with the version of the database.
// const version = 1

// // Migrate migrates the Postgres schema to the current version.
// func validateSchema(db *sql.DB) error {
// 	sourceInstance, err := iofs.New(fs, "migrations")
// 	if err != nil {
// 		return err
// 	}
// 	targetInstance, err := postgres.WithInstace(db, new(postgres.Config))
// 	if err != nil {
// 		return err
// 	}
// 	m, err := migrate.NewWithInstace("iofs", sourceInstance, "postgres", targetInstance)
// 	if err != nil {
// 		return err
// 	}
// 	err = m.Migrate(version) // current version
// 	if err != nil && err != migrate.ErrNoChange {
// 		return err
// 	}
// 	return sourceInstance.Close()
// // }
