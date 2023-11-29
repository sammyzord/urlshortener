package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sammyzord/urlshortener/internal"
)

func main() {
	internal.ConnectDB()

	internal.DB.AutoMigrate(&internal.URL{})

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", internal.HomepageHandler)
	router.POST("/", internal.UrlPostHandler)
	router.Run()
}
