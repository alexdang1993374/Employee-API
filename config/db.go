package config

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func Connect() *bun.DB {
	dsn := "postgres://postgres:postgres@localhost:5432/shorten_url?sslmode=disable"
	pgconn := sql.OpenDB(
		pgdriver.NewConnector(
			pgdriver.WithDSN(dsn),
			pgdriver.WithAddr("localhost:5432"),
			pgdriver.WithUser("postgres"),
			pgdriver.WithPassword("CC1993374"),
			pgdriver.WithDatabase("employees"),
		),
	)
	db := bun.NewDB(pgconn, pgdialect.New())

	return db
}
