package main

import (
	"FeedAggregator/internal/database"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (c config) handlerPostFeedFollow(writer http.ResponseWriter, request *http.Request, dbUsr database.User) {
	type requestBody struct {
		FeedId string `json:"feed_id"`
	}

	body := requestBody{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&body)
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, "Couldn't decode parameters")
		fmt.Println("Error decoding parameter: " + err.Error())
		return
	}
	if len(body.FeedId) == 0 {
		respondWithError(writer, http.StatusBadRequest, "Invalid request body")
		return
	}

	uUID, err := uuid.NewUUID()
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, "Internal server error")
		fmt.Println("Error creating uuid: " + err.Error())
		return
	}

	feedFollow, err := c.DB.CreateFeedFollow(request.Context(), database.CreateFeedFollowParams{
		ID:        uUID.String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    dbUsr.ID,
		FeedID:    body.FeedId,
	})
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, "Internal server error")
		fmt.Println("Error creating feedFollow: " + err.Error())
		return
	}

	respondWithJSON(writer, http.StatusOK, databaseFeedFollowTo(feedFollow))
}
