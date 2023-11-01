package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(writer http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Panicln("Responding with 5** error: ", message)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(writer, code, errResponse{
		Error: message,
	})
}

func respondWithJson(writer http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v\n", payload)
		writer.WriteHeader(500)
		return
	}
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(data)
}