-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    balance NUMERIC(15,2) NOT NULL DEFAULT 0.00
);

-- +goose Down
DROP TABLE users;
