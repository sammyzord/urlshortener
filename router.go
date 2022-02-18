package main

import (
	"net/http"
	controllers "urlshortener/controllers"
)

func router() {
	http.HandleFunc("/new", controllers.NewUrl)
}
