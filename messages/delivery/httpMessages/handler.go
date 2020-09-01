package httpMessages

import (
	"ChatAPI/messages"
	"ChatAPI/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
)

type GetStructChatID struct {
	ChatId string `json:"chat" binding:"required"`
}

type HandlerMessage struct {
	UseCase messages.UseCaseMessages
}

func NewHandlerMessage(useCase messages.UseCaseMessages) *HandlerMessage {
	return &HandlerMessage{
		UseCase: useCase,
	}
}

func (h *HandlerMessage) AddMessage(c *gin.Context) {
	var json models.Message

	// unmarshal
	err := c.ShouldBindJSON(&json)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, "invalid json")
	}

	// validation data
	if json.Author == "" || json.Chat == "" || strings.TrimSpace(json.Text) == "" || json.Id != 0 {
		c.String(http.StatusBadRequest, "invalid json")
	}

	_, err = strconv.Atoi(json.Author)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid Author")
	}

	_, err = strconv.Atoi(json.Chat)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid Chat")
	}

	// add message -> UseCase
	message, err := h.UseCase.AddMessage(c, &json)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, message)
}

func (h *HandlerMessage) GetMessagesByIdHandler(c *gin.Context) {
	var json GetStructChatID

	// unmarshal
	err := c.ShouldBindJSON(&json)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, "invalid json")
	}

	// validation data
	if strings.TrimSpace(json.ChatId) == "" {
		log.Error(err)
		c.String(http.StatusBadRequest, "invalid json")
	}

	intChatID, err := strconv.Atoi(json.ChatId)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, "invalid json")
	}

	// add message -> UseCase
	message, err := h.UseCase.GetMessagesByChatID(c, &models.Chat{
		Id: intChatID,
	})
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, message)
}
