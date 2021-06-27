package core

import (
	"fmt"

	"github.com/filipeandrade6/vigia-go/internal/database"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/repository"
	"github.com/joho/godotenv"
)

func Main() error {

	err := godotenv.Load()
	if err != nil {
		return err
	}

	// logger,  := zap.NewProduction()
	// defer logger.Sync()

	dbCfg := database.NewConfig()
	p, err := database.NewPool(dbCfg)
	if err != nil {
		return err
	}

	cam, err := repository.GetByID(p, "1")
	if err != nil {
		return err
	}

	fmt.Println(cam)

	return nil
}
