package main

import (
	"fmt"
	"os"

	"github.com/filipeandrade6/vigia-go/internal/gerencia/core"
	"github.com/filipeandrade6/vigia-go/internal/logger"
)

func main() {
	log, err := logger.New("GERENCIA")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	if err := core.Run(log); err != nil {
		log.Errorw("startup", "ERROR", err)
		os.Exit(1)
	}
}
