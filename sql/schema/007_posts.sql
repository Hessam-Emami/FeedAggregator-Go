-- +goose Up
CREATE TABLE posts
(
    id           varchar primary key,
    created_at   timestamp,
    updated_at   timestamp,
    title        varchar        not null,
    url          varchar unique not null,
    description  varchar,
    published_at timestamp,
    feed_id      varchar        not null
        references feeds (id) on delete CASCADE
);

-- +goose Down
DROP TABLE posts;