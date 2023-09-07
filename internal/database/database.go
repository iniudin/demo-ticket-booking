package database

import (
	"context"
	"fmt"
	"github.com/exaring/otelpgx"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Begin(ctx context.Context) pgx.Tx {
	tx, err := Pool.Begin(ctx)
	if err != nil {
		panic(err)
	}
	return tx
}

func Commit(ctx context.Context, tx pgx.Tx) {
	if err := tx.Commit(ctx); err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			panic(rollbackErr)
		}
		panic(err)
	}
}

func SetupDatabase(ctx context.Context) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Panic(err)
	}

	config.ConnConfig.Tracer = otelpgx.NewTracer()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Panic(err)
	}
	Pool = pool
}
