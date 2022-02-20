package models

import (
	"database/sql"
	"log"
	"urlshortener/database"
	"urlshortener/utils"

	"github.com/speps/go-hashids/v2"
)

type Url struct {
	Id   int    `json:"id"`
	Url  string `json:"url"`
	Hash string `json:"hash"`
}

func Create(url string) string {
	query := "INSERT INTO urls (url) VALUES ($1) RETURNING id;"
	var lastId int

	err := database.DB.QueryRow(query, url).Scan(&lastId)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	hd := hashids.NewData()
	hd.Salt = utils.Salt
	h, _ := hashids.NewWithData(hd)
	hash, _ := h.Encode([]int{lastId})

	updateQuery := "UPDATE urls SET hash = $1 WHERE id = $2"

	_, err = database.DB.Exec(updateQuery, hash, lastId)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return hash
}

func Retrieve(hash string) (string, error) {
	var url string
	err := database.DB.QueryRow("SELECT url FROM urls WHERE hash = $1", hash).Scan(&url)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}
		return "", err
	}
	return url, nil
}
