package model

import (
	"swordsman/logger"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Conn *gorm.DB

func init() {
	dsn := "root:12345678@tcp(192.168.15.51:3306)/civ_1839?charset=utf8&parseTime=True&loc=Local"
	// dsn := "3b9rcbawlkld:pscale_pw_7wiOR6xAsF26eNYv46PVDZlj1oLl2lTt-WLHdACBoPI@tcp(l4vq442qa5b8.ap-southeast-2.psdb.cloud)/wg_1001?tls=true"
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
}
