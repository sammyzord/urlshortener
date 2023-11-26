package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomepageHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func UrlPostHandler(context *gin.Context) {
	url := context.PostForm("url")
	context.HTML(http.StatusOK, "new-url.tmpl", gin.H{
		"newUrl": url,
	})
}
