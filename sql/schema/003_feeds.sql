-- +goose Up
CREATE TABLE feeds
(
    id         varchar primary key,
    created_at timestamp      not null,
    updated_at timestamp      not null,
    name       varchar        not null,
    url        varchar unique not null,
    user_id    varchar        not null REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;