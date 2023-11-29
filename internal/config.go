package internal

import "os"

func GetBaseUrl() string {
	var BASE_URL string = os.Getenv("BASE_URL")
	if BASE_URL == "" {
		return "http://localhost:8080"
	} else {
		return BASE_URL
	}
}
