package controller

import (
	"net/http"
	"swordsman/consts"
	"swordsman/middleware"
	"swordsman/msg"
	"swordsman/services"
	"swordsman/utils"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResultT(consts.CodeOk, "", nil))
}

func GetRole(c *gin.Context) {
	account := middleware.JwtAccount(c)
	user := services.UserService.GetRole(account.ID)
	c.JSON(http.StatusOK, utils.ResultT(consts.CodeOk, "", gin.H{
		"account": account,
		"role":    user.ToFront(),
	}))
}

func CreateRole(c *gin.Context) {
	var createRole msg.User
	_ = c.Bind(&createRole)
	if createRole == (msg.User{}) {
		c.JSON(http.StatusOK, utils.ResultT(consts.CodeErrorParam, "", nil))
		return
	}
	if err := services.UserService.CreateRole(createRole.ID, createRole.Name); err != nil {
		c.JSON(http.StatusOK, utils.ResultT(consts.CodeErrorParam, "", nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResultT(consts.CodeOk, "", nil))
}
