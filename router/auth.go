package router

import (
	"swordsman/controller"
	"swordsman/middleware"

	"github.com/gin-gonic/gin"
)

func loginRouter(r *gin.RouterGroup) {
	g := r.Group("/auth")
	// 登录
	g.POST("login", middleware.AuthMiddleware.LoginHandler)
	// 注册
	g.PUT("login", controller.Register)
	// 登出
	g.POST("logout", middleware.AuthMiddleware.LogoutHandler)
}
