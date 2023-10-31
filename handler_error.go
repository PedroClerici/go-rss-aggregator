package main

import "net/http"

func handlerError(writer http.ResponseWriter, req *http.Request) {
	respondWithError(writer, 400, "Something went wrong!")
}