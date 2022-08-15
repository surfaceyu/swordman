package services

import (
	"swordsman/model"
	"swordsman/utils"
	"sync"
	"time"
)

var (
	// 所有的业务类都在这儿声明
	UserService *UserMgr
)

type UserMgr struct {
	serverId int32
	users    sync.Map
}

func Init(serverId int32) {
	NewUserMgr(serverId).Run()
	ChatService.Init()
}

func NewUserMgr(serverId int32) *UserMgr {
	UserService = &UserMgr{
		serverId: serverId,
	}
	return UserService
}

func (um *UserMgr) Run() {
	utils.GoSafe(func() {
		go um.lru()
	})
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
	return *um.GetRoleByUid(uid)
}

func (um *UserMgr) GetRoleByUid(uid int64) *model.User {
	if uid <= 0 {
		return &model.User{}
	}
	// 优先从缓存中读取玩家信息
	user, ok := um.getRole(uid)
	if ok {
		return user
	}
	// 找不到玩家在从数据库中读取信息
	user = model.GetRole(uid)
	user.TTL = time.Now().Add(3600 * time.Second)
	um.storeRole(user)
	return user
}

// 读取缓存中的玩家信息
func (um *UserMgr) getRole(uid int64) (*model.User, bool) {
	u, got := um.users.Load(uid)
	user, ok := u.(*model.User)
	return user, got && ok
}

// 缓存玩家信息
func (um *UserMgr) storeRole(user *model.User) {
	um.users.Store(user.ID, user)
}

// 过期玩家信息
func (um *UserMgr) lru() {
	var uidLru []int64
	um.users.Range(func(id, u any) bool {
		uid := id.(int64)
		user := u.(*model.User)
		if user.TTL.Before(time.Now()) {
			uidLru = append(uidLru, uid)
		}
		return true
	})
	for _, uid := range uidLru {
		um.users.Delete(uid)
	}
}
