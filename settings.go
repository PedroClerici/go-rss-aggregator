package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type settings struct {
	port string	
	databaseURL string
}

func getSettings() settings {
	// Read your env file(s) and load them into ENV for this process.
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT variable not found in the environment")
	}

	databaseURL := os.Getenv("DB_URL")
	if databaseURL == "" {
		log.Fatal("DB_URL variable not found in the environment")
	}

	return settings {
		port: port,
		databaseURL: databaseURL,
	}
}