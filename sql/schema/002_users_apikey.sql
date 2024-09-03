-- +goose Up
--ALTER TABLE users ADD COLUMN apikey VARCHAR(64) UNIQUE NOT NULL DEFAULT (
--    encode(sha256(random()::text::bytea), 'hex')
--);

ALTER TABLE users ADD COLUMN apikey VARCHAR(64);

UPDATE users
SET apikey = lower(hex(randomblob(32)))
WHERE apikey IS NULL;

CREATE UNIQUE INDEX idx_users_apikey ON users(apikey);

-- +goose Down
ALTER TABLE users
    DROP COLUMN apikey;
