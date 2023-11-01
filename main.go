package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/PedroClerici/go-rss-aggregator/internal/database"
	"github.com/PedroClerici/go-rss-aggregator/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	_ "github.com/lib/pq"
)

func main() {
	settings := getSettings()

	// Opens a connection to database.
	conn, err := sql.Open("postgres", settings.databaseURL)
	if err != nil {
		log.Fatal("Unable to connect to database")
	}

	db := database.New(conn)

	router := chi.NewRouter()

	// Middlewares setup
	router.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://*", "http://*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"*"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300,
  }))
	router.Use(middleware.Logger)

	server := &http.Server{
		Handler: router,
		Addr: ":" + settings.port,
	}

	router.Mount("/", routes.StatusResource{}.Routes())
	router.Mount("/users", routes.UsersResource{DB: db}.Routes())

	log.Printf("Server starting on port: %s\n", settings.port)

	if server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}