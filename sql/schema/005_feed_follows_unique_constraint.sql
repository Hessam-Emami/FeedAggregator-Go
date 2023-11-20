-- +goose Up
DELETE
FROM feed_follows
WHERE (feed_id, user_id) IN (SELECT feed_id, user_id
                             FROM feed_follows
                             GROUP BY feed_id, user_id
                             HAVING COUNT(*) > 1)
  AND ctid NOT IN (SELECT MIN(ctid)
                   FROM feed_follows
                   GROUP BY feed_id, user_id
                   HAVING COUNT(*) > 1);
ALTER TABLE feed_follows
    ADD CONSTRAINT unique_feed_user_combination UNIQUE (feed_id, user_id);

-- +goose Down
ALTER TABLE feed_follows
    DROP CONSTRAINT IF EXISTS unique_feed_user_combination;