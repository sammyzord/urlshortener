package main

import (
	"net/http"
	controllers "urlshortener/controllers"
)

func router() {
	http.HandleFunc("/", controllers.GetUrl)
	http.HandleFunc("/new", controllers.NewUrl)
}
