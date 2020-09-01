package postgresql

import (
	"ChatAPI/models"
	"errors"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
)

type PostgresqlStorage struct {
	DB *gorm.DB
}

func NewPSQLRepository(db *gorm.DB) *PostgresqlStorage {
	return &PostgresqlStorage{
		DB: db,
	}
}

func (r *PostgresqlStorage) AddUserToDB(c *gin.Context, user *models.User) (*models.User, error) {
	ormUser := models.UsersGORM{
		Username:  user.Username,
		CreatedAt: time.Now(),
	}

	//// because with an ordinary error, the ID of the next user will increase by the number of previous errors
	//db := r.DB.Find(&ormUser, "username = ?", ormUser.Username)
	//if db.Error == nil {
	//	return nil, errors.New("Username is already exist!")
	//}

	// create user
	db := r.DB.Create(&ormUser)
	if db.Error != nil {
		return nil, errors.New("Can not add user!")
	}

	return &models.User{
		Id:        ormUser.Id,
		Username:  ormUser.Username,
		CreatedAt: ormUser.CreatedAt,
	}, nil
}
