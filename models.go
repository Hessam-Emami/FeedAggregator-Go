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
