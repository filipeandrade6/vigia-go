package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	DBUser         string
	DBPass         string
	DBHost         string
	DBPort         int
	DBName         string
	DBPoolMaxConns int
}

func (c *Config) getDSN() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s pool_max_conns=%d",
		c.DBUser,
		c.DBPass,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBPoolMaxConns,
	)
}

func NewConfig() *Config {
	var cfg Config

	cfg.DBUser = os.Getenv("DATABASE_USER")
	cfg.DBPass = os.Getenv("DATABASE_PASS")
	cfg.DBHost = os.Getenv("DATABASE_HOST")

	var err error
	cfg.DBPort, err = strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Fatalln("Error on load env var:", err.Error())
	}

	cfg.DBName = os.Getenv("DATABASE_NAME")

	cfg.DBPoolMaxConns, err = strconv.Atoi(os.Getenv("DATABASE_POOLMAXCONNS"))
	if err != nil {
		log.Fatalln("Error on load env var", err.Error())
	}

	return &cfg
}
