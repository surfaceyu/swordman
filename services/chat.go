package services

import (
	"fmt"
	"swordsman/consts"
	"swordsman/model"
	"sync"
	"time"
)

var (
	maxMsgLen int = 20
	// 所有的业务类都在这儿声明
	ChatService = &ChatMgr{
		serverId: 1001,
	}
)

type ChatMgr struct {
	serverId int32
	Channels sync.Map
}

func (cm *ChatMgr) Init() {
	// 创建世界频道
	cm.CreateChannel(consts.ChannelWorld, 0, 0)
	// 创建地图频道
	cm.CreateChannel(consts.ChannelMap, 1001, 0)
	// 创建帮派频道
	cm.CreateChannel(consts.ChannelGuild, 1001, 0)
	// 加载旧数据
	chats := model.GetChatByType(consts.ChannelWorld)
	for _, v := range chats {
		chName := cm.ChannelName(v.Typ, v.From, v.To)
		ch, ok := cm.GetChannel(chName)
		if !ok {
			continue
		}
		ch.Push(ChatMessage{
			Uid:       v.From,
			Channel:   chName,
			Message:   v.Data,
			CreatedAt: v.CreatedAt,
		})
	}
}

func (cm *ChatMgr) ChannelName(typ string, fromId int64, toId int64) string {
	switch typ {
	case consts.ChannelWorld:
		fromId = 0
		toId = 0
	case consts.ChannelMap, consts.ChannelGuild:
		toId = 0
	case consts.ChannelUser:
		// no data change
	}
	return fmt.Sprintf("%s%d%d", typ, fromId, toId)
}

func (cm *ChatMgr) GetChannel(channelName string) (*ChatChannel, bool) {
	ch, got := cm.Channels.Load(channelName)
	channel, ok := ch.(*ChatChannel)
	return channel, got && ok
}

func (cm *ChatMgr) CreateChannel(typ string, fromId int64, toId int64) *ChatChannel {
	ch := &ChatChannel{
		Type:     typ,
		FromId:   fromId,
		ToId:     toId,
		Messages: []ChatMessage{},
	}
	cm.Channels.Store(ch.ChannelName(), ch)
	return ch
}

func (cm *ChatMgr) SendMsg(typ string, from int64, to int64, data string) {
	chName := cm.ChannelName(typ, from, to)
	ch, ok := cm.GetChannel(chName)
	if !ok {
		ch = cm.CreateChannel(typ, from, to)
	}
	ch.AddChat(from, data)
}

func (cm *ChatMgr) GetMsg(channelName string, limit int) []ChatMessage {
	ch, ok := cm.GetChannel(channelName)
	if !ok {
		return []ChatMessage{}
	}
	return ch.GetMsg(limit)
}

type ChatChannel struct {
	Type     string
	FromId   int64
	ToId     int64
	Messages []ChatMessage
}

func (cc *ChatChannel) ChannelName() string {
	return ChatService.ChannelName(cc.Type, cc.FromId, cc.ToId)
}

func (cc *ChatChannel) Push(msg ChatMessage) {
	cc.Messages = append(cc.Messages, msg)
	msgLen := len(cc.Messages)
	if msgLen > maxMsgLen {
		cc.Messages = cc.Messages[msgLen-maxMsgLen:]
	}
}

func (cc *ChatChannel) AddChat(uid int64, data string) {
	msg := ChatMessage{
		Uid:       uid,
		Channel:   cc.ChannelName(),
		Message:   data,
		CreatedAt: time.Time{},
	}
	cc.Push(msg)
	msg.model(*cc).Save()
}

func (cc *ChatChannel) GetMsg(count int) []ChatMessage {
	msgLen := len(cc.Messages)
	if count >= msgLen {
		count = msgLen
	}
	return cc.Messages[msgLen-count:]
}

type ChatMessage struct {
	Uid       int64
	Channel   string
	Message   string
	CreatedAt time.Time
}

func (cm *ChatMessage) model(cc ChatChannel) *model.Chat {
	fromUser := UserService.GetRoleByUid(cm.Uid)
	toUser := UserService.GetRoleByUid(cc.ToId)
	return &model.Chat{
		Typ:      cc.Type,
		From:     cm.Uid,
		FromName: fromUser.UserName,
		Data:     cm.Message,
		To:       cc.ToId,
		ToName:   toUser.UserName,
	}
}
