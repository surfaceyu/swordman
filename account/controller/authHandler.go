package controller

import (
	"net/http"
	"swordsman/consts"
	"swordsman/model"
	"swordsman/msg"
	"swordsman/services"
	"swordsman/utils"

	"github.com/gin-gonic/gin"
)

// 创建角色接口
func Register(c *gin.Context) {
	var createAccount msg.Account
	_ = c.Bind(&createAccount)
	if "" == createAccount.Account || "" == createAccount.Passwd {
		c.JSON(http.StatusOK, utils.ResultT(consts.CodeErrorParam, "", nil))
		return
	}
	// 判断账号是否存在
	var accountGot = services.UserService.FindUser(createAccount.Account)
	if accountGot.ID != 0 {
		c.JSON(http.StatusOK, utils.ResultT(consts.CodeErrorHasReg, "", nil))
		return
	}
	// 创建新账号
	err := services.UserService.SaveUser(&model.Account{
		Account: createAccount.Account,
		Passwd:  createAccount.Passwd,
	})
	if err != nil {
		c.JSON(http.StatusOK, utils.ResultT(consts.CodeErrorParam, "", nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResultT(consts.CodeOk, "", nil))
}

func GetServer(c *gin.Context) {
	servers := model.FindServer()
	c.JSON(http.StatusOK, utils.ResultT(consts.CodeOk, "", servers))
}
