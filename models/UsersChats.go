package models

type UsersChatsGORM struct {
	Id   int `gorm:"AUTO_INCREMENT;column:users_chats_id"`
	User int `gorm:"column:user_id"`
	Chat int `gorm:"column:chat_id"`
}

// Set name table
func (UsersChatsGORM) TableName() string {
	return "users_chats"
}
