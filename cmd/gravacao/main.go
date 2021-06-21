package main

import (
	"context"
	"fmt"
	"os"

	dbConfig "github.com/filipeandrade6/vigia-go/internal/database"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

type Camera struct {
	ID             int
	IP             string
	Descricao      string
	Porta          int
	Canal          int
	UsuarioCamera  string
	SenhaCamera    string
	Cidade         string
	Geolocalizacao string
	Marca          string
	Modelo         string
	Informacao     string
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	fmt.Printf("Acessing %s ...", dbConfig.DbName)

	// ! utilizar pgxpool e conectar com pgxpool.Connect() - concurrency safe
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName))
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	test := &dbConfig.Camera{}

	err = conn.QueryRow(context.Background(), "SELECT porta, descricao FROM camera WHERE id=$1", 1).Scan(&test.Porta, &test.Descricao)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(test.Descricao))

	return nil
}
