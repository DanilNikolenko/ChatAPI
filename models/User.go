package models

import "time"

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type UsersGORM struct {
	Id        int       `gorm:"column:id"`
	Username  string    `gorm:"column:username"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

// Set name table
func (UsersGORM) TableName() string {
	return "users"
}
