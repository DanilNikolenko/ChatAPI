package messages

import (
	"ChatAPI/models"

	"github.com/gin-gonic/gin"
)

type UseCaseMessages interface {
	AddMessage(c *gin.Context, message *models.Message) (int, error)
	GetMessagesByChatID(c *gin.Context, chat *models.Chat) (*[]models.Message, error)
}
