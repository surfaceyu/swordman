package router

import (
	"swordsman/game/controller"
	"swordsman/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	g := r.Group("/api")
	g.Use(middleware.AuthMiddleware.MiddlewareFunc())
	userRouter(g)
	chatRouter(g)
}

func userRouter(r *gin.RouterGroup) {
	g := r.Group("/user")
	g.GET("role", controller.GetRole)
	g.PUT("role", controller.CreateRole)
}

func chatRouter(r *gin.RouterGroup) {
	g := r.Group("/chat")
	g.GET("chat", controller.GetChatMsg)
	g.POST("chat", controller.SendChatMsg)
}
