
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE Users_Chats
(
    Users_Chats_ID SERIAL PRIMARY KEY,
    User_ID INTEGER NOT NULL,
    Chat_ID INTEGER NOT NULL,

    FOREIGN KEY (User_ID) REFERENCES Users (Id) ON DELETE CASCADE,
    FOREIGN KEY (Chat_ID) REFERENCES Chats (Id) ON DELETE CASCADE
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE Users_Chats;
