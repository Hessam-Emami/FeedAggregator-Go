package main

import (
	"FeedAggregator/internal/auth"
	"FeedAggregator/internal/database"
	"net/http"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (c config) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		apiKey, err := auth.GetAuthHeader(request)
		if err != nil {
			respondWithError(writer, http.StatusUnauthorized, "Not authorised")
			return
		}

		user, err := c.DB.GetUserByApiKey(request.Context(), apiKey)
		if err != nil {
			respondWithError(writer, http.StatusUnauthorized, "Not authorised")
			return
		}

		handler(writer, request, user)
	}
}
