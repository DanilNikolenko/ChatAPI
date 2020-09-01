package usecase

import (
	"ChatAPI/models"
	"ChatAPI/users"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type UserUseCase struct {
	UserStorage users.UsersRepository
}

func NewUsersUseCase(ur users.UsersRepository) *UserUseCase {
	return &UserUseCase{
		UserStorage: ur,
	}
}

func (r *UserUseCase) AddUser(c *gin.Context, user *models.User) (int, error) {
	myUser, err := r.UserStorage.AddUserToDB(c, user)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return myUser.Id, nil
}
