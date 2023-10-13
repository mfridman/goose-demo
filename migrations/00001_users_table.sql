-- +goose Up
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username text NOT NULL
);

-- +goose Down
DROP TABLE users;
