package main

import (
	"FeedAggregator/internal/database"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (c config) handlerDeleteFeedFollow(writer http.ResponseWriter, request *http.Request, dbUsr database.User) {
	ffId := chi.URLParam(request, "FeedFollowId")
	if len(ffId) == 0 {
		respondWithError(writer, http.StatusBadRequest, "Feed Follow is required")
		return
	}

	isDeleted, err := c.DB.DeleteFeedFollow(request.Context(), ffId)
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, "Internal server error")
		fmt.Println("Error creating feedFollow: " + err.Error())
		return
	}
	if isDeleted {
		respondWithJSON(writer, http.StatusOK, "Feed deleted: "+ffId)
	} else {
		respondWithError(writer, http.StatusInternalServerError, "Could not delete feed : "+ffId)
	}
}

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

func (c config) handlerGetFeedFollow(writer http.ResponseWriter, request *http.Request, dbUsr database.User) {
	feedFollowsDb, err := c.DB.GetFeedFollowsByUserId(request.Context(), dbUsr.ID)
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, "Internal server error")
		fmt.Println("Error getting feedFollow: " + err.Error())
		return
	}
	feedFollows := make([]FeedFollowDto, 0)
	for _, ff := range feedFollowsDb {
		feedFollows = append(feedFollows, databaseFeedFollowTo(ff))
	}
	respondWithJSON(writer, http.StatusOK, feedFollows)
}
