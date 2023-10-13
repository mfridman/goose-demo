-- +goose Up
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username text
);

-- +goose Down
DROP TABLE users;
