package router

import (
	"swordsman/controller"
	"swordsman/middleware"

	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.RouterGroup) {
	g := r.Group("/user")
	g.Use(middleware.AuthMiddleware.MiddlewareFunc())
	g.GET("hello", controller.Hello)
	g.GET("role", controller.GetRole)
	g.PUT("role", controller.CreateRole)
}
