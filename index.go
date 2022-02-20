package main

import (
	"log"
	"net/http"
	"urlshortener/database"
	"urlshortener/utils"
)

func main() {
	database.Connect()
	router()
	log.Fatal(http.ListenAndServe(":"+utils.Port, nil))
}
