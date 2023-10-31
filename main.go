package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Read your env file(s) and load them into ENV for this process.
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT variable not found in the environment")
	}

	router := chi.NewRouter()

	// CORS set up
	router.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://*", "http://*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"*"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300,
  }))

	server := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	router.Get("/health", handlerReadiness)
	router.Get("/error", handlerError)

	log.Printf("Server starting on port: %s\n", portString)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}