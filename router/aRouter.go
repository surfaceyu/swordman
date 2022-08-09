package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	g := r.Group("/api")
	loginRouter(g)
	userRouter(g)
}
