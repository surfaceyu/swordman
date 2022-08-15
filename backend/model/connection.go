package model

import (
	"swordsman/config"
	"swordsman/logger"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var (
	Conn *gorm.DB
)

func init() {
	config.Init()
	dsn := viper.GetString("mysql")
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		logger.Fatal(err)
	}
	db.SetLogger(logger.GetLogger())
	db.LogMode(true)
	db.DB().SetMaxOpenConns(2000)
	db.DB().SetMaxIdleConns(2000)
	db.DB().SetConnMaxLifetime(time.Second * 300)
	if err = db.DB().Ping(); err != nil {
		logger.Fatal(err)
	}
	Conn = db

	autoMigrate()
}

func autoMigrate() {
	Conn.AutoMigrate(&Account{})
	Conn.AutoMigrate(&User{})
	Conn.AutoMigrate(&Cache{})
	Conn.AutoMigrate(&Chat{})
	Conn.AutoMigrate(&Server{})
}
