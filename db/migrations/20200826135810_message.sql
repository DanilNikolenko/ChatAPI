
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE Messages
(
    Id SERIAL PRIMARY KEY,
    Chat_Id INTEGER NOT NULL,
    Author_Id INTEGER NOT NULL,
    Text VARCHAR NOT NULL,

    Created_at TIMESTAMP,

    FOREIGN KEY (Author_Id) REFERENCES Users (Id) ON DELETE CASCADE,
    FOREIGN KEY (Chat_Id) REFERENCES Chats (Id) ON DELETE CASCADE
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE Messages;
