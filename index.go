package main

import (
	"log"
	"net/http"
	"urlshortener/database"
)

func main() {
	database.Connect()
	router()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
