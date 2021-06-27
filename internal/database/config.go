package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	dbUser         string
	dbPass         string
	dbHost         string
	dbPort         int
	dbName         string
	dbPoolMaxConns int
	dbDSN          string
}

func NewConfig() *Config {
	var cfg Config

	cfg.dbUser = os.Getenv("DATABASE_USER")
	cfg.dbPass = os.Getenv("DATABASE_PASS")
	cfg.dbHost = os.Getenv("DATABASE_HOST")

	var err error
	cfg.dbPort, err = strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Fatalln("Error on load env var:", err.Error())
	}

	cfg.dbName = os.Getenv("DATABASE_NAME")

	cfg.dbPoolMaxConns, err = strconv.Atoi(os.Getenv("DATABASE_POOLMAXCONNS"))
	if err != nil {
		log.Fatalln("Error on load env var", err.Error())
	}

	cfg.dbDSN = fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s pool_max_conns=%d",
		cfg.dbUser,
		cfg.dbPass,
		cfg.dbHost,
		cfg.dbPort,
		cfg.dbName,
		cfg.dbPoolMaxConns,
	)

	return &cfg
}
