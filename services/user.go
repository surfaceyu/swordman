package services

import "swordsman/model"

var (
	// 所有的业务类都在这儿声明
	UserService = &UserMgr{}
)

type UserMgr struct {
}

func (p *UserMgr) FindUser(id string) model.Account {
	return model.FindUser(id)
}

func (p *UserMgr) SaveUser(user *model.Account) error {
	return model.AddUser(user)
}

func (p *UserMgr) CreateRole() {

}
