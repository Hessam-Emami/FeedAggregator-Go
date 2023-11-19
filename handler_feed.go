package main

import (
	"FeedAggregator/internal/database"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (c config) handlerPostFeed(writer http.ResponseWriter, request *http.Request, dbUsr database.User) {
	type requestBody struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	body := requestBody{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&body)
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, "Couldn't decode parameters")
		fmt.Println("Error decoding parameter: " + err.Error())
		return
	}
	if len(body.Name) == 0 || len(body.URL) == 0 {
		respondWithError(writer, http.StatusBadRequest, "Invalid request body")
		return
	}

	uUID, err := uuid.NewUUID()
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, "Internal server error")
		fmt.Println("Error creating uuid: " + err.Error())
		return
	}

	feed, err := c.DB.CreateFeed(request.Context(), database.CreateFeedParams{
		ID:        uUID.String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      body.Name,
		Url:       body.URL,
		UserID:    dbUsr.ID,
	})
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, "Internal server error")
		fmt.Println("Error creating feed: " + err.Error())
		return
	}

	respondWithJSON(writer, http.StatusOK, databaseFeedToFeedDto(feed))
}
