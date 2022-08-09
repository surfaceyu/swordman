package utils

import (
	"fmt"
	"swordsman/consts"

	"github.com/gin-gonic/gin"
)

func ResultT(code consts.CodeError, msg string, data any) gin.H {
	return gin.H{
		"code":    code.Code,
		"message": fmt.Sprintf("%s%s", msg, code.Message),
		"data":    data,
	}
}
