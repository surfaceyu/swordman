package model

import (
	"swordsman/msg"
	"time"
)

type Server struct {
	ID        int16 `gorm:"primary_key"`
	Name      string
	Host      string
	Port      int16
	Status    int8
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 指定表名
func (Server) TableName() string {
	return "server"
}

func (s Server) ToFront() msg.Server {
	return msg.Server{
		ID:   s.ID,
		Name: s.Name,
		Host: s.Host,
		Port: s.Port,
	}
}

func FindServer() []Server {
	var s []Server
	Conn.Where("Status < ?", 8).Find(&s)
	return s
}
