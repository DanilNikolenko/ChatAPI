package httpChats

import (
	"ChatAPI/chats"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpointsChats(c *gin.Engine, uc chats.UseCaseChats) {
	h := NewHandlerChats(uc)

	// create group
	urs := c.Group("/chats")
	{
		// add in group
		urs.POST("/add", h.AddChat)
		urs.POST("/get", h.GetChatsHandler)
	}

	fmt.Println("Methods chats registered!")
}
