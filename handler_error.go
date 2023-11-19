package main

import "net/http"

func handlerError(writer http.ResponseWriter, request *http.Request) {
	respondWithError(writer, 500, "Interval Server Error")
}
