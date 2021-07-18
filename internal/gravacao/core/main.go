package core

import (
	"github.com/filipeandrade6/vigia-go/internal/database"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/client"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/server"

	"github.com/joho/godotenv"
)

func Main() error {

	err := godotenv.Load()
	if err != nil {
		return err
	}

	// logger,  := zap.NewProduction()
	// defer logger.Sync()

	server.StartServer("tcp", "localhost:12346")
	a := client.GetGerenciaClient("localhost:12347")

	dbCfg := a.GetDatabase()

	p, err := database.NewPool(dbCfg)
	if err != nil {
		return err
	}

	return nil
}
