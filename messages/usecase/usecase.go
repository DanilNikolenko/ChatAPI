package usecase

import (
	"ChatAPI/messages"
	"ChatAPI/models"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type MessageUseCase struct {
	MessageStorage messages.MessageRepository
}

func NewMessagesUseCase(mr messages.MessageRepository) *MessageUseCase {
	return &MessageUseCase{
		MessageStorage: mr,
	}
}

func (r *MessageUseCase) AddMessage(c *gin.Context, message *models.Message) (int, error) {
	myUser, err := r.MessageStorage.AddMessageToDB(c, message)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return myUser.Id, nil
}

func (r *MessageUseCase) GetMessagesByChatID(c *gin.Context, chat *models.Chat) (*[]models.Message, error) {
	myMessages, err := r.MessageStorage.GetMessagesByChatID(c, chat)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return myMessages, nil
}
