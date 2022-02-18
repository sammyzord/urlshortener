package database

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("pgx", "postgresql://user:password@localhost:5432/urlshortener")
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}

}
