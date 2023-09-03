package Getter

import (
	"com.github.goscaffold/pkg/dbs"
	"com.github.goscaffold/pkg/models/UserModel"
	"com.github.goscaffold/pkg/result"
	"fmt"
)

var UserGetter IUserGetter

func init() {
	UserGetter = NewIUserGetterImpl()
}

type IUserGetter interface {
	GetUserList() []*UserModel.UserModelImpl
	GetUserDetail(id int) *result.ErrorResult
	CreateUser(user *UserModel.UserModelImpl) *result.ErrorResult
	UpdateUser(id int, user *UserModel.UserModelImpl) *result.ErrorResult
	DeleteUser(id int) *result.ErrorResult
}

type IUserGetterImpl struct {
}

func NewIUserGetterImpl() *IUserGetterImpl {
	return &IUserGetterImpl{}
}

// 创建用户
func (this *IUserGetterImpl) CreateUser(user *UserModel.UserModelImpl) *result.ErrorResult {
	db := dbs.Orm.Create(user)
	if db.Error != nil {
		return result.Result(nil, db.Error)
	}
	return result.Result(user, nil)
}

// 获取用户列表
func (this *IUserGetterImpl) GetUserList() (users []*UserModel.UserModelImpl) {
	dbs.Orm.Find(&users)
	return users
}

// 获取用户详情
func (this *IUserGetterImpl) GetUserDetail(id int) *result.ErrorResult {

	user := UserModel.New()
	db := dbs.Orm.Where("id=?", id).First(user)

	if db.Error != nil || db.RowsAffected == 0 {
		return result.Result(nil, fmt.Errorf("not found user, id=%d", id))
	}

	return result.Result(user, nil)
}

// 更新用户
func (this *IUserGetterImpl) UpdateUser(id int, user *UserModel.UserModelImpl) *result.ErrorResult {
	db := dbs.Orm.Where("id=?", id).Updates(user)
	if db.Error != nil {
		return result.Result(nil, db.Error)
	}
	return result.Result(user, nil)
}

// 删除用户
func (this *IUserGetterImpl) DeleteUser(id int) *result.ErrorResult {
	user := UserModel.New()
	db := dbs.Orm.Where("id=?", id).Delete(user)
	if db.Error != nil {
		return result.Result(nil, db.Error)
	}
	return result.Result(user, nil)
}
