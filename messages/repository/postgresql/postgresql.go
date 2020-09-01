package postgresql

import (
	"ChatAPI/models"
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PostgresqlStorage struct {
	DB *gorm.DB
}

func NewMessagesPSQLRepository(db *gorm.DB) *PostgresqlStorage {
	return &PostgresqlStorage{
		DB: db,
	}
}

func (r *PostgresqlStorage) AddMessageToDB(c *gin.Context, message *models.Message) (*models.Message, error) {
	intIDUser, _ := strconv.Atoi(message.Author)
	intIDChat, _ := strconv.Atoi(message.Chat)

	//// find user
	//db := r.DB.Find(&models.UsersGORM{}, "id = ?", intIDuser)
	//if db.Error != nil {
	//	return nil, errors.New("User is not exist ")
	//}
	//
	//// find chat
	//db = r.DB.Find(&models.ChatsGORM{}, "Id = ?", intIDchat)
	//if db.Error != nil {
	//	return nil, errors.New("Chat is not exist ")
	//}

	// find user in chat
	db := r.DB.Find(&models.UsersChatsGORM{}, "user_id = ? AND chat_id = ?", intIDUser, intIDChat)
	if db.Error != nil {
		return nil, errors.New("User do not exist in this chat ")
	}

	ormMessages := models.MessagesGORM{
		Chat:      intIDChat,
		Author:    intIDUser,
		Text:      message.Text,
		CreatedAt: time.Now(),
	}

	// create message
	db = r.DB.Create(&ormMessages)
	if db.Error != nil {
		return nil, errors.New("Can not create message ")
	}

	return &models.Message{
		Id:        ormMessages.Id,
		Chat:      message.Chat,
		Author:    message.Author,
		Text:      ormMessages.Text,
		CreatedAt: ormMessages.CreatedAt,
	}, nil
}

func (r *PostgresqlStorage) GetMessagesByChatID(c *gin.Context, chat *models.Chat) (*[]models.Message, error) {
	messageORM := []models.MessagesGORM{}

	// find and sort messages
	db := r.DB.Order("created_at").Find(&messageORM, "chat_id = ?", chat.Id)
	if db.Error != nil {
		return nil, errors.New("Error find messages ")
	}

	// check len message slice
	if len(messageORM) == 0 {
		return &[]models.Message{}, nil
	}

	// convert to models message
	rez := []models.Message{}
	for _, message := range messageORM {
		chatSTR := strconv.Itoa(message.Chat)
		authorSTR := strconv.Itoa(message.Author)

		rez = append(rez, models.Message{
			Id:        message.Id,
			Chat:      chatSTR,
			Author:    authorSTR,
			Text:      message.Text,
			CreatedAt: message.CreatedAt,
		})
	}

	return &rez, nil
}
