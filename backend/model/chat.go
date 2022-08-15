package model

import (
	"time"
)

type Chat struct {
	ID        uint `gorm:"primary_key"`
	Typ       string
	From      int64
	FromName  string
	To        int64
	ToName    string
	Data      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// 指定表名
func (Chat) TableName() string {
	return "chat"
}

func GetChatAll() []Chat {
	var chats []Chat
	Conn.Find(&chats)
	return chats
}

func GetChatByType(typ string) []Chat {
	var chats []Chat
	Conn.Where("typ = ?", typ).Find(&chats)
	return chats
}

func GetUserChatByType(typ string, uid int64) []Chat {
	var chats []Chat
	Conn.Where("typ = ? and from = ?", typ, uid).Or("typ = ? and to = ?", typ, uid).Find(&chats)
	return chats
}

func (c *Chat) Save() {
	Conn.Save(&c)
}
