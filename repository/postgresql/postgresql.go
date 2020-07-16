package postgresql

import (
	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

// New returns database instance used in this package
func New(connString string) (*sqlx.DB, error) {
	config, err := pgx.ParseConfig(connString)
	if err != nil {
		panic(err)
	}

	db := sqlx.NewDb(stdlib.OpenDB(*config), "pgx")
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
