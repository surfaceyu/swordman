package router

import (
	"swordsman/account/controller"
	"swordsman/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	g := r.Group("/api")
	loginRouter(g)
}

func loginRouter(r *gin.RouterGroup) {
	g := r.Group("/auth")
	// 登录
	g.POST("login", middleware.AuthMiddleware.LoginHandler)
	// 注册
	g.PUT("login", controller.Register)
	// 登出
	g.POST("logout", middleware.AuthMiddleware.LogoutHandler)
	// 获取服务器列表
	g.GET("server", controller.GetServer)
}
