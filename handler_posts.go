package main

import (
	"FeedAggregator/internal/database"
	"fmt"
	"net/http"
	"strconv"
)

func (c config) handlerGetPost(writer http.ResponseWriter, request *http.Request, dbUsr database.User) {
	limitStr := request.URL.Query().Get("limit")
	limit := 10
	if specifiedLimit, err := strconv.Atoi(limitStr); err == nil {
		limit = specifiedLimit
	}

	posts, err := c.DB.GetPostsByUser(request.Context(), database.GetPostsByUserParams{
		UserID: dbUsr.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, "Internal server error")
		fmt.Println("Error creating feed: " + err.Error())
		return
	}

	postsDto := make([]PostDto, 0)
	for _, postDb := range posts {
		postsDto = append(postsDto, databasePostToPostDto(postDb))
	}
	respondWithJSON(writer, http.StatusOK, postsDto)
}
