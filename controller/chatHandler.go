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

func SendChatMsg(c *gin.Context) {
	account := middleware.JwtAccount(c)
	uid := services.UserService.GetUid(account.ID)

	var chatMsg msg.SendChatMsg
	_ = c.Bind(&chatMsg)

	services.ChatService.SendMsg(chatMsg.Channel, uid, chatMsg.To, chatMsg.Data)
	c.JSON(http.StatusOK, utils.ResultT(consts.CodeOk, "", nil))
}

func GetChatMsg(c *gin.Context) {
	account := middleware.JwtAccount(c)
	uid := services.UserService.GetUid(account.ID)

	var chatMsg msg.GetChatMsg
	_ = c.Bind(&chatMsg)
	limit := 10
	if chatMsg.Limit < limit {
		limit = chatMsg.Limit
	}
	channelName := services.ChatService.ChannelName(chatMsg.Channel, uid, chatMsg.To)
	c.JSON(http.StatusOK, utils.ResultT(consts.CodeOk, "", services.ChatService.GetMsg(channelName, limit)))
}
