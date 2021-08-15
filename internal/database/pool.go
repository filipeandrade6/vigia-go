package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPool(cfg *Config) (*pgxpool.Pool, error) {
	poolCfg, err := pgxpool.ParseConfig(cfg.getDSN())
	if err != nil {
		return nil, fmt.Errorf("error loading connection pool database config: %w", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), poolCfg)
	if err != nil {
		return nil, fmt.Errorf("error creating connection pool: %w", err)
	}

	return pool, nil
}

// simples abaixo
// https://github.com/MarioCarrion/videos/blob/c11c3ae9c7ad52c544e5ca88062aca79342ec61d/2021/02/03-go-database-postgresql-part-1/postgresql_pgx.go

// completo abaixo
// https://github.com/johanbrandhorst/bazel-mono/blob/1150f6f2280417261fcdc4205245c678c6faac0c/cmd/go-server/users/users.go

// outro abaixo
// https://github.com/manniwood/iidy/blob/b30b8d79cb911c3c213975e4c0e167ed939546bb/pgstore/pgstore.go
