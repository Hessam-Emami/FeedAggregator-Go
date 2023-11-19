-- +goose Up
CREATE TABLE feed_follows
(
    id         varchar primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    feed_id    varchar   not null REFERENCES feeds (id) ON DELETE CASCADE,
    user_id    varchar   not null REFERENCES users (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feed_follows;