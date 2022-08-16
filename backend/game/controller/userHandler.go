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

func GetRole(c *gin.Context) {
	account := middleware.JwtAccount(c)
	user := services.UserService.GetRoleByAccount(account.ID)
	c.JSON(http.StatusOK, utils.ResultT(consts.CodeOk, "", gin.H{
		"account": account,
		"role":    user.ToFront(),
	}))
}

func CreateRole(c *gin.Context) {
	var createRole msg.CreateRole
	_ = c.Bind(&createRole)
	account := middleware.JwtAccount(c)
	if createRole == (msg.CreateRole{}) {
		c.JSON(http.StatusOK, utils.ResultT(consts.CodeErrorParam, "", nil))
		return
	}
	if len(createRole.Name) > 21 {
		c.JSON(http.StatusOK, utils.ResultT(consts.CodeErrorNameTooLong, "", nil))
		return
	}
	if err := services.UserService.CreateRole(account.ID, createRole.Name); err != nil {
		c.JSON(http.StatusOK, utils.ResultT(consts.CodeErrorParam, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResultT(consts.CodeOk, "", nil))
}
