package utils

import "os"

func getEnv(key string, fallback string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	} else {
		return val
	}
}

var Port string = getEnv("PORT", "8080")

var BaseUrl string = getEnv("BASE_URL", "localhost:"+Port)

var Salt string = getEnv("SALT", "this is my salt")

var DatabaseUrl string = getEnv("DATABASE_URL", "postgresql://user:password@localhost:5432/urlshortener")
