package httpMessages

import (
	"ChatAPI/messages"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpointsMessages(c *gin.Engine, uc messages.UseCaseMessages) {
	h := NewHandlerMessage(uc)

	// create group
	urs := c.Group("/messages")
	{
		// add in group
		urs.POST("/add", h.AddMessage)
		urs.POST("/get", h.GetMessagesByIdHandler)
	}

	fmt.Println("Methods messages registered!")
}
