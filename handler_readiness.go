package main

import "net/http"

func handlerReadiness(writer http.ResponseWriter, req *http.Request) {
	respondWithJson(writer, 200, struct{}{})
}