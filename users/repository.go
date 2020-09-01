package users

import (
	"ChatAPI/models"

	"github.com/gin-gonic/gin"
)

type UsersRepository interface {
	AddUserToDB(c *gin.Context, user *models.User) (*models.User, error)
}
