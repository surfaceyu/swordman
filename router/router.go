package router

import (
	"swordsman/controller"
	"swordsman/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	g := r.Group("/api")
	loginRouter(g)
	userRouter(g)
	chatRouter(g)
}

func loginRouter(r *gin.RouterGroup) {
	g := r.Group("/auth")
	// 登录
	g.POST("login", middleware.AuthMiddleware.LoginHandler)
	// 注册
	g.PUT("login", controller.Register)
	// 登出
	g.POST("logout", middleware.AuthMiddleware.LogoutHandler)
}

func userRouter(r *gin.RouterGroup) {
	g := r.Group("/user")
	g.Use(middleware.AuthMiddleware.MiddlewareFunc())
	g.GET("hello", controller.Hello)
	g.GET("role", controller.GetRole)
	g.PUT("role", controller.CreateRole)
}

func chatRouter(r *gin.RouterGroup) {
	g := r.Group("/chat")
	g.Use(middleware.AuthMiddleware.MiddlewareFunc())
	g.GET("chat", controller.GetChatMsg)
	g.POST("chat", controller.SendChatMsg)
}
