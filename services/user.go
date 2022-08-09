package services

import (
	"swordsman/model"
)

var (
	// 所有的业务类都在这儿声明
	UserService = &UserMgr{
		serverId: 1001,
	}
)

type UserMgr struct {
	serverId int32
}

func (um *UserMgr) IncrUid(accountId uint) int64 {
	return int64(um.serverId*1000000) + int64(accountId)
}

func (um *UserMgr) FindUser(account string) model.Account {
	return model.FindUser(account)
}

func (um *UserMgr) SaveUser(user *model.Account) error {
	return model.AddUser(user)
}

func (um *UserMgr) CreateRole(accountId uint, name string) error {
	return model.CreateRole(&model.User{
		ID:       um.IncrUid(accountId),
		UserName: name,
	})
}

func (um *UserMgr) GetRole(accountId uint) model.User {
	uid := um.IncrUid(accountId)
	return model.GetRole(uid)
}
