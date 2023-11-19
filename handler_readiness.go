package main

import "net/http"

func handlerReadiness(writer http.ResponseWriter, request *http.Request) {
	respondWithJSON(writer, 200, struct {
		Status string `json:"status"`
	}{Status: "ok"})
}
