package httpUsers

import (
	"ChatAPI/users"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpointsUsers(c *gin.Engine, us users.UseCaseUsers) {
	h := NewHandlerUsers(us)

	//c.POST("/users/add", h.AddUser)

	urs := c.Group("/users")
	{
		urs.POST("/add", h.AddUser)
	}

	fmt.Println("Methods users registered!")
}
