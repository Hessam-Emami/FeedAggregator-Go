// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, created_at, updated_at, feed_id, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at, feed_id, user_id
`

type CreateFeedFollowParams struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedID    string
	UserID    string
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.FeedID,
		arg.UserID,
	)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FeedID,
		&i.UserID,
	)
	return i, err
}
