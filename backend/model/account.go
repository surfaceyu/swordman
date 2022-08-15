package model

import (
	"swordsman/consts"
	"swordsman/msg"
	"time"
)

type Account struct {
	ID        uint   `gorm:"primary_key"`
	Account   string `gorm:"uniqueIndex"`
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
		ID:      a.ID,
		Account: a.Account,
		Passwd:  a.Passwd,
	}
}

func FindUser(account string) Account {
	var loginAccount Account
	Conn.First(&loginAccount, Account{Account: account})
	return loginAccount
}

func AddUser(user *Account) error {
	if err := Conn.Create(user).Error; err != nil {
		return consts.ADD_DATA_ERROR
	}
	return nil
}
