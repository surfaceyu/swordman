package model

import (
	"swordsman/consts"
	"swordsman/logger"
	"swordsman/msg"
	"time"
)

type User struct {
	ID        string `gorm:"primary_key;size:128"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// 指定表名
func (User) TableName() string {
	return "user"
}

func (u User) ToFront() msg.User {
	return msg.User{
		ID:   u.ID,
		Name: u.Name,
	}
}

func GetRole(id string) User {
	var role User
	Conn.First(&role, User{ID: id})
	return role
}

func CreateRole(user *User) error {
	if err := Conn.Create(user).Error; err != nil {
		logger.Fatal(err)
		return consts.ADD_DATA_ERROR
	}
	return nil
}
