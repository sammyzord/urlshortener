package models

import (
	"log"
	"urlshortener/database"

	"github.com/speps/go-hashids/v2"
)

type Url struct {
	Id   int    `json:"id"`
	Url  string `json:"url"`
	Hash string `json:"hash"`
}

func Create(url string) {
	query := "INSERT INTO urls (url) VALUES ($1) RETURNING id;"
	var lastId int

	err := database.DB.QueryRow(query, url).Scan(&lastId)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	h, _ := hashids.NewWithData(hd)
	hash, _ := h.Encode([]int{lastId})

	updateQuery := "UPDATE urls SET hash = $1 WHERE id = $2"

	_, err = database.DB.Exec(updateQuery, hash, lastId)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func Retrieve(hash string) {

}
