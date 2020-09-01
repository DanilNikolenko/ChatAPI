package messages

import (
	"ChatAPI/models"

	"github.com/gin-gonic/gin"
)

type MessageRepository interface {
	AddMessageToDB(c *gin.Context, message *models.Message) (*models.Message, error)
	GetMessagesByChatID(c *gin.Context, chat *models.Chat) (*[]models.Message, error)
}
