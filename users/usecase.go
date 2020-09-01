package users

import (
	"ChatAPI/models"

	"github.com/gin-gonic/gin"
)

type UseCaseUsers interface {
	AddUser(c *gin.Context, user *models.User) (int, error)
}
