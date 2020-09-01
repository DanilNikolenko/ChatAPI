package postgresql

import (
	"ChatAPI/models"
	"errors"
	"sort"
	"strconv"
	"time"

	"github.com/prometheus/common/log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PostgresqlStorageChats struct {
	DB *gorm.DB
}

func NewPSQLRepositoryChats(db *gorm.DB) *PostgresqlStorageChats {
	return &PostgresqlStorageChats{
		DB: db,
	}
}

func (r *PostgresqlStorageChats) AddChatToDB(c *gin.Context, chat *models.Chat) (*models.Chat, error) {
	// convert slice str to int
	// make slice pq.Int64Array
	intSliceUsers := make([]int, len(chat.Users), len(chat.Users))

	// cycle to convert
	for index, user := range chat.Users {
		tempUserID, err := strconv.Atoi(user)
		if err != nil {
			log.Error(err)
			return nil, errors.New("Invalid slice users ")
		}
		intSliceUsers[index] = tempUserID
	}

	// make orm struct
	ormChat := models.ChatsGORM{
		Name:      chat.Name,
		CreatedAt: time.Now(),
	}

	// because with an ordinary error, the ID of the next user will increase by the number of previous errors
	db := r.DB.Find(&models.ChatsGORM{}, "name = ?", ormChat.Name)
	if db.Error == nil {
		return nil, errors.New("Name is already exist ")
	}

	// START TRANSACTION ===========
	tx := r.DB.Begin()

	// create chat
	db = tx.Create(&ormChat)
	if db.Error != nil {
		tx.Rollback()
		return nil, errors.New("Can not add chat!")
	}

	// created connection Users_Chats
	err := AddUsersChats(tx, &intSliceUsers, ormChat.Id)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("Can not add chat, user error!")
	}

	if tx.Commit().Error != nil {
		return nil, errors.New("Can not add chat!")
	}
	// END TRANSACTION ============

	return &models.Chat{
		Id:        ormChat.Id,
		Name:      ormChat.Name,
		Users:     chat.Users, // return string
		CreatedAt: ormChat.CreatedAt,
	}, nil
}

func AddUsersChats(tx *gorm.DB, users *[]int, chatID int) error {

	// add users_chats to db
	for _, user := range *users {
		if tx.Create(&models.UsersChatsGORM{
			User: user,
			Chat: chatID,
		}).Error != nil {
			return errors.New("Can not add UsersChats")
		}
	}

	return nil
}

func (r *PostgresqlStorageChats) GetChatsByUserID(c *gin.Context, user *models.User) (*[]models.Chat, error) {
	UsersChats := []models.UsersChatsGORM{}

	// find user in chats
	db := r.DB.Find(&UsersChats, "user_id = ?", user.Id)
	if db.Error != nil {
		return nil, errors.New("Error find chats ")
	}

	// check len chats
	if len(UsersChats) == 0 {
		return &[]models.Chat{}, nil
	}

	// find chats range
	rez := []models.Chat{}

	for _, uniqueChat := range UsersChats {

		// find chats
		tempChatORM := models.ChatsGORM{}
		db := r.DB.Find(&tempChatORM, "id = ?", uniqueChat.Chat)
		if db.Error != nil {
			log.Error(db.Error)
			return nil, errors.New("Error: chat does not exist ")
		}

		// add to rez slice one chat
		rez = append(rez, models.Chat{
			Id:   tempChatORM.Id,
			Name: tempChatORM.Name,
			//Users:     UserSTR,
			CreatedAt: tempChatORM.CreatedAt,
		})
	}

	// --------sort by last message
	// make map time last message in chat
	ChatMap := map[int]time.Time{}

	// add to map time last message
	for _, chat := range rez {
		TempMessage := models.MessagesGORM{}

		// find last message
		db := r.DB.Last(&TempMessage, "chat_id = ?", chat.Id)
		if db.Error != nil {
			log.Error(db.Error)
		}

		// add to map time
		ChatMap[chat.Id] = TempMessage.CreatedAt
	}

	// sorting a slice using maps
	sort.Slice(rez, func(current, next int) bool {
		return ChatMap[rez[current].Id].After(ChatMap[rez[next].Id])
	})
	// ----------

	return &rez, nil
}
