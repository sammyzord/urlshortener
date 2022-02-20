package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"urlshortener/models"
	"urlshortener/utils"
)

type UrlJson struct {
	Url string `json:"url"`
}

func NewUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		var url UrlJson

		err := json.NewDecoder(r.Body).Decode(&url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		hash := models.Create(url.Url)

		redirectLink := utils.BaseUrl + "/" + hash

		json.NewEncoder(w).Encode(redirectLink)
	default:
		http.Error(w, "Sorry, only POST method is supported.", http.StatusMethodNotAllowed)
	}

}

func GetUrl(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		path := strings.Split(r.URL.Path, "/")

		if len(path) == 2 && path[1] != "" {
			hash := path[1]

			url, err := models.Retrieve(hash)
			if err != nil {
				http.Error(w, "Url code not found.", http.StatusNotFound)
			} else {
				if !strings.Contains(url, "http://") && !strings.Contains(url, "https://") {
					url = "http://" + url
				}
				http.Redirect(w, r, url, http.StatusSeeOther)
			}

		} else {
			http.Error(w, "Please inform a valid address.", http.StatusBadRequest)
		}

	default:
		http.Error(w, "Sorry, only GET method is supported.", http.StatusMethodNotAllowed)
	}
}
