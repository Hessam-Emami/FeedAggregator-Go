-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostsByUser :many
SELECT posts.*
FROM posts
         JOIN feeds f on f.id = posts.feed_id
WHERE f.user_id = $1 LIMIT $2;
