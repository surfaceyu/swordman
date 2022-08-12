package main

import (
	"swordsman/game/router"
	"swordsman/logger"
	_ "swordsman/logger"
	"swordsman/middleware"
	"swordsman/model"
	"swordsman/services"

	"github.com/gin-gonic/gin"
	flag "github.com/spf13/pflag"
	godeamon "github.com/surfaceyu/godeamon"
)

func main() {
	serverId := flag.Int32P("server", "s", 1001, "game serve id")

	godeamon.App.Run()

	defer model.Conn.Close()

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(middleware.Cors()) //使用中间件
	router.InitRouter(r)

	services.Init(*serverId)

	logger.Info("服务器启动中 [:4567]。。。。。。")
	if err := r.Run(":4567"); err != nil {
		logger.Fatalf("[%s]端口被占用,请修改配置", "4567")
	}
}
