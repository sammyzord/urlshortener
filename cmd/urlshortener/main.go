package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sammyzord/urlshortener/internal"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", internal.HomepageHandler)
	router.POST("/", internal.UrlPostHandler)
	router.Run()
}
