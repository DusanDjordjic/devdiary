package db

import (
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB = nil

func Connect() error {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres dbname=devdiary password=postgres sslmode=disable")
	if err != nil {
		return err
	}

	DB = db

	return DB.Ping()
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}

	return nil
}
