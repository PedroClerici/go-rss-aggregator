package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PedroClerici/go-rss-aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UsersResource struct{
	DB *database.Queries
}

func (rs UsersResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", rs.Create)

	return r
}

func (rs *UsersResource) Create(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json"name"`
	}

	// Decoding parameters
	d := json.NewDecoder(r.Body)
	params := parameters{}
	if err := d.Decode(&params); err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}

	// Creating user on database
	user, err := rs.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondWithJson(w, 200, user)
}