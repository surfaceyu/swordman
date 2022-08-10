package services

import (
	"swordsman/model"
	"sync"
)

var (
	// 所有的业务类都在这儿声明
	UserService = &UserMgr{
		serverId: 1001,
	}
)

type UserMgr struct {
	serverId int32
	users    sync.Map
}

func (um *UserMgr) GetUid(accountId uint) int64 {
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
		ID:       um.GetUid(accountId),
		UserName: name,
	})
}

func (um *UserMgr) GetRoleByAccount(accountId uint) model.User {
	uid := um.GetUid(accountId)
	return um.GetRoleByUid(uid)
}

func (um *UserMgr) GetRoleByUid(uid int64) model.User {
	if uid <= 0 {
		return model.User{}
	}
	// 优先从缓存中读取玩家信息
	user, ok := um.getRole(uid)
	if ok {
		return user
	}
	// 找不到玩家在从数据库中读取信息
	user = model.GetRole(uid)
	return user
}

// 读取缓存中的玩家信息
func (um *UserMgr) getRole(uid int64) (model.User, bool) {
	u, got := um.users.Load(uid)
	user, ok := u.(model.User)
	return user, got && ok
}

// 缓存玩家信息
func (um *UserMgr) storeRole(user model.User) {
	um.users.Store(user.ID, user)
}

// 过期玩家信息
func (um *UserMgr) lru() {
}
