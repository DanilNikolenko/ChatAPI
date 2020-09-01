package chats

import (
	"ChatAPI/models"

	"github.com/gin-gonic/gin"
)

type ChatsRepository interface {
	AddChatToDB(c *gin.Context, chat *models.Chat) (*models.Chat, error)
	GetChatsByUserID(c *gin.Context, user *models.User) (*[]models.Chat, error)
}
