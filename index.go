package main

import (
	"fmt"
	"log"
	"net/http"
	"urlshortener/database"
	"urlshortener/utils"
)

func main() {
	database.Connect()
	router()
	fmt.Println("Serving at port", utils.Port)
	log.Fatal(http.ListenAndServe(":"+utils.Port, nil))
}
