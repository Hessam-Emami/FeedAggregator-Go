package main

import (
	"FeedAggregator/internal/database"
	"time"
)

type UserDto struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUserDto(dbUsr database.User) UserDto {
	return UserDto{
		Id:        dbUsr.ID,
		CreatedAt: dbUsr.CreatedAt,
		UpdatedAt: dbUsr.UpdatedAt,
		Name:      dbUsr.Name,
		ApiKey:    dbUsr.ApiKey,
	}
}

type FeedDto struct {
	Id            string     `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Name          string     `json:"name"`
	Url           string     `json:"url"`
	UserId        string     `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func databaseFeedToFeedDto(dbFeed database.Feed) FeedDto {
	var lastFetchedAt *time.Time
	if dbFeed.LastFetchedAt.Valid {
		lastFetchedAt = &dbFeed.LastFetchedAt.Time
	} else {
		lastFetchedAt = nil
	}
	return FeedDto{
		Id:            dbFeed.ID,
		CreatedAt:     dbFeed.CreatedAt,
		UpdatedAt:     dbFeed.UpdatedAt,
		Name:          dbFeed.Name,
		Url:           dbFeed.Url,
		UserId:        dbFeed.UserID,
		LastFetchedAt: lastFetchedAt,
	}
}

type FeedFollowDto struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedID    string    `json:"feed_id"`
	UserID    string    `json:"user_id"`
}

func databaseFeedFollowTo(dbFeed database.FeedFollow) FeedFollowDto {
	return FeedFollowDto{
		Id:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		UserID:    dbFeed.UserID,
		FeedID:    dbFeed.FeedID,
	}
}
