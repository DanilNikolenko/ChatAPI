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
	var json models.Chat

	// unmarshal
	err := c.ShouldBindJSON(&json)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, "invalid json")
	}

	// validation data
	if len(json.Users) <= 1 || strings.TrimSpace(json.Name) == "" || json.Id != 0 {
		log.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	}

	// add chat -> UseCase
	chat, err := h.UseCase.AddChat(c, &json)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, chat)
}

func (h *HandlerChats) GetChatsHandler(c *gin.Context) {
	var json PostGetChats

	// unmarshal
	err := c.ShouldBindJSON(&json)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, "invalid json")
	}

	// validation data
	if strings.TrimSpace(json.UserID) == "" {
		log.Error(err)
		c.String(http.StatusBadRequest, "invalid json")
	}

	intUserID, err := strconv.Atoi(json.UserID)
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
