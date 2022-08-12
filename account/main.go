package main

import (
	"swordsman/account/router"
	"swordsman/logger"
	_ "swordsman/logger"
	"swordsman/middleware"
	"swordsman/model"

	"github.com/gin-gonic/gin"
	godeamon "github.com/surfaceyu/godeamon"
)

func main() {
	godeamon.App.Run()

	defer model.Conn.Close()

	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	r.Use(middleware.Cors()) //使用中间件
	router.InitRouter(r)

	r.Static("/game", "./public")
	r.Static("/assets", "./public/assets")
	r.StaticFile("/vite.svg", "./public/vite.svg")

	logger.Info("服务器启动中 [:4566]。。。。。。")
	if err := r.Run(":4566"); err != nil {
		logger.Fatalf("[%s]端口被占用,请修改配置", "4566")
	}
}
