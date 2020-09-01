package httpChats

import (
	"ChatAPI/chats"
	"ChatAPI/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
)

type PostGetChats struct {
	UserID string `json:"user" binding:"required"`
}

type HandlerChats struct {
	UseCase chats.UseCaseChats
}

func NewHandlerChats(useCase chats.UseCaseChats) *HandlerChats {
	return &HandlerChats{
		UseCase: useCase,
	}
}

func (h *HandlerChats) AddChat(c *gin.Context) {
	var jsonData models.Chat

	// unmarshal
	err := c.ShouldBindJSON(&jsonData)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, "invalid json")
	}

	// validation data
	if len(jsonData.Users) <= 1 || strings.TrimSpace(jsonData.Name) == "" || jsonData.Id != 0 {
		log.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	}

	// add chat -> UseCase
	chat, err := h.UseCase.AddChat(c, &jsonData)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, chat)
}

func (h *HandlerChats) GetChatsHandler(c *gin.Context) {
	var jsonData PostGetChats

	// unmarshal
	err := c.ShouldBindJSON(&jsonData)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, "invalid json")
	}

	// validation data
	if strings.TrimSpace(jsonData.UserID) == "" {
		log.Error(err)
		c.String(http.StatusBadRequest, "invalid json")
	}

	intUserID, err := strconv.Atoi(jsonData.UserID)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, "invalid json")
	}

	// add chat -> UseCase
	chat, err := h.UseCase.GetChats(c, &models.User{
		Id: intUserID,
	})
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, chat)
}
