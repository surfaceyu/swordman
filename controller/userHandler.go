package controller

import (
	"net/http"
	"wordGame/consts"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, consts.Code(consts.CodeOk))
}
