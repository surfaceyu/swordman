package model

import (
	"swordsman/consts"
	"swordsman/logger"
	"swordsman/msg"
	"time"
)

type Account struct {
	ID        string `gorm:"primary_key;size:128"`
	Passwd    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// 指定表名
func (Account) TableName() string {
	return "account"
}

func (a Account) ToFront() msg.Account {
	return msg.Account{
		ID:     a.ID,
		Passwd: a.Passwd,
	}
}

func FindUser(id string) Account {
	var loginAccount Account
	Conn.First(&loginAccount, Account{ID: id})
	return loginAccount
}

func AddUser(user *Account) error {
	if err := Conn.Create(user).Error; err != nil {
		logger.Fatal(err)
		return consts.ADD_DATA_ERROR
	}
	return nil
}
