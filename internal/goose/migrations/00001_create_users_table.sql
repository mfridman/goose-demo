-- +goose Up
CREATE TABLE users (
    id int NOT NULL PRIMARY KEY,
    username text
);

INSERT INTO users VALUES
(0, 'root'),
(1, 'goosey');

-- +goose Down
DROP TABLE users;