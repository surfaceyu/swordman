package router

import (
	"swordsman/controller"
	"swordsman/middleware"

	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.RouterGroup) {
	g := r.Group("/hello")
	g.Use(middleware.AuthMiddleware.MiddlewareFunc())
	g.GET("hello", controller.Hello)
}
