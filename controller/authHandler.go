package controller

import (
	"net/http"
	"swordsman/consts"
	"swordsman/model"
	"swordsman/services"
	"swordsman/utils"

	"github.com/gin-gonic/gin"
)

// 创建角色接口
func Register(c *gin.Context) {
	var createAccount model.Account
	_ = c.Bind(&createAccount)
	if "" == createAccount.ID || "" == createAccount.Passwd {
		c.JSON(http.StatusOK, utils.ResultT(consts.CodeErrorParam, "", nil))
		return
	}
	// 判断账号是否存在
	var accountGot = services.UserService.FindUser(createAccount.ID)
	if accountGot.ID != "" {
		c.JSON(http.StatusOK, utils.ResultT(consts.CodeErrorHasReg, "", nil))
		return
	}
	// 创建新账号
	err := services.UserService.SaveUser(&createAccount)
	if err != nil {
		c.JSON(http.StatusOK, utils.ResultT(consts.CodeErrorParam, "", nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResultT(consts.CodeOk, "", nil))
}
