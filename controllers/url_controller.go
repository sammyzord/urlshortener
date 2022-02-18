package controllers

import (
	"encoding/json"
	"net/http"
	"urlshortener/models"
)

type UrlJson struct {
	Url string `json:"url"`
}

func NewUrl(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var url UrlJson

		err := json.NewDecoder(r.Body).Decode(&url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		models.Create(url.Url)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(url.Url)
	default:
		http.Error(w, "Sorry, only POST method is supported.", http.StatusMethodNotAllowed)
		return
	}

}

func GetUrl(w http.ResponseWriter, r *http.Request) {}
