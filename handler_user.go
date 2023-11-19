package main

import (
	"FeedAggregator/internal/auth"
	"FeedAggregator/internal/database"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (c config) handlerGetUserByApiKey(writer http.ResponseWriter, request *http.Request) {
	apiKey, err := auth.GetAuthHeader(request)
	if err != nil {
		respondWithError(writer, http.StatusBadRequest, "Error getting authorisation header")
		return
	}

	usr, err := c.DB.GetUserByApiKey(request.Context(), apiKey)
	if err != nil {
		respondWithError(writer, http.StatusNotFound, "Couldn't find the user")
		return
	}

	respondWithJSON(writer, http.StatusOK, databaseUserToUserDto(usr))
}

func (c config) handlerPostUser(writer http.ResponseWriter, request *http.Request) {
	type requestBody struct {
		Name string `json:"name"`
	}

	body := requestBody{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&body)
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, "Couldn't decode parameters")
		fmt.Println("Error decoding parameter: " + err.Error())
		return
	}
	if len(body.Name) == 0 {
		respondWithError(writer, http.StatusBadRequest, "Invalid request body")
		return
	}

	uUID, err := uuid.NewUUID()
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, "Internal server error")
		fmt.Println("Error creating uuid: " + err.Error())
		return
	}

	usr, err := c.DB.CreateUser(request.Context(), database.CreateUserParams{
		ID:        uUID.String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      body.Name,
	})
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, "Internal server error")
		fmt.Println("Error creating user: " + err.Error())
		return
	}

	respondWithJSON(writer, http.StatusOK, databaseUserToUserDto(usr))
}
