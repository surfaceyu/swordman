package model

import (
	"swordsman/consts"
	"swordsman/msg"
	"time"
)

type User struct {
	ID        int64 `gorm:"primary_key"`
	UserName  string
	TTL       time.Time `gorm:"-"`
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
		Name: u.UserName,
	}
}

func GetRole(id int64) *User {
	var role = &User{}
	Conn.First(role, User{ID: id})
	return role
}

func CreateRole(user *User) error {
	if err := Conn.Create(user).Error; err != nil {
		return consts.ADD_DATA_ERROR
	}
	return nil
}

type Cache struct {
	ID         string `gorm:"primary_key;size:128"`
	DataString string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}

func (Cache) TableName() string {
	return "cache"
}

func (c *Cache) Get() *Cache {
	Conn.First(&c, Cache{ID: c.ID})
	return c
}

func (c *Cache) Save() {
	Conn.Save(&c)
}
