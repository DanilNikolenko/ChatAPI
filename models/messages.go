package models

import "time"

type Message struct {
	Id        int       `json:"id"`
	Chat      string    `json:"chat" binding:"required"`
	Author    string    `json:"author" binding:"required"`
	Text      string    `json:"text" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type MessagesGORM struct {
	Id        int       `gorm:"column:id"`
	Chat      int       `gorm:"column:chat_id"`
	Author    int       `gorm:"column:author_id"`
	Text      string    `gorm:"column:text"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

// Set name table
func (MessagesGORM) TableName() string {
	return "messages"
}
