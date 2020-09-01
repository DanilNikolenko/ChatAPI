package usecase

import (
	"ChatAPI/chats"
	"ChatAPI/models"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type ChatrUseCase struct {
	ChatsStorage chats.ChatsRepository
}

func NewChatsUseCase(chr chats.ChatsRepository) *ChatrUseCase {
	return &ChatrUseCase{
		ChatsStorage: chr,
	}
}

func (r *ChatrUseCase) AddChat(c *gin.Context, chat *models.Chat) (int, error) {
	myChat, err := r.ChatsStorage.AddChatToDB(c, chat)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return myChat.Id, nil
}

func (r *ChatrUseCase) GetChats(c *gin.Context, user *models.User) (*[]models.Chat, error) {
	myChats, err := r.ChatsStorage.GetChatsByUserID(c, user)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return myChats, nil
}
