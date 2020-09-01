package models

import (
	"time"
)

type Chat struct {
	Id        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Users     []string  `json:"users,omitempty" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type ChatsGORM struct {
	Id        int       `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

// Set name table
func (ChatsGORM) TableName() string {
	return "chats"
}
