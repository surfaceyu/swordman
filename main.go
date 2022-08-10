package main

import (
	"swordsman/logger"
	_ "swordsman/logger"
	"swordsman/middleware"
	"swordsman/model"
	"swordsman/router"
	"swordsman/services"

	"github.com/gin-gonic/gin"
)

func main() {
	defer model.Conn.Close()

	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	r.Use(middleware.Cors()) //使用中间件
	router.InitRouter(r)

	r.Static("/game", "./public")
	r.Static("/assets", "./public/assets")
	r.StaticFile("/vite.svg", "./public/vite.svg")

	services.ChatService.Init()

	logger.Info("服务器启动中。。。。。。")
	if err := r.Run(":4567"); err != nil {
		logger.Fatalf("[%s]端口被占用,请修改配置", "4567")
	}
}
