
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE Chats
(
    Id SERIAL PRIMARY KEY,
    Name VARCHAR(30) UNIQUE NOT NULL,

    Created_at TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE Chats;
