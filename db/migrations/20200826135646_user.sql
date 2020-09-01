
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE Users
(
    Id SERIAL PRIMARY KEY,
    Username VARCHAR(30) UNIQUE NOT NULL,

    Created_at TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE Users;