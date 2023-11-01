package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type StatusResource struct{}

func (rs StatusResource) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", rs.handlerHealth)
	router.Get("/panic", rs.handlerError)

	return router
}

func (rs StatusResource) handlerHealth(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, "Hello, world!")
}

func (rs StatusResource) handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong!")
}