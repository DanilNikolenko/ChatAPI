package chats

import (
	"ChatAPI/models"

	"github.com/gin-gonic/gin"
)

type UseCaseChats interface {
	AddChat(c *gin.Context, chat *models.Chat) (int, error)
	GetChats(c *gin.Context, user *models.User) (*[]models.Chat, error)
}
