package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomepageHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func UrlPostHandler(context *gin.Context) {
	base_url := GetBaseUrl()
	url := context.PostForm("url")

	newUrl, err := shortenUrl(url)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "new-url.tmpl", gin.H{
			"newUrl": "Error generating URL",
		})
		return
	}

	context.HTML(http.StatusOK, "new-url.tmpl", gin.H{
		"newUrl": base_url + "/" + newUrl,
	})
}
