package database

import (
	"database/sql"
	"log"
	"urlshortener/utils"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("pgx", utils.DatabaseUrl)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}

}
