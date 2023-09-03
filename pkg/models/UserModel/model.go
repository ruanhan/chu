package UserModel

import "golang.org/x/crypto/bcrypt"

// 创建 Users struct
type UserModelImpl struct {
	// Gorm：主键 自增
	Id int `json:"id" gorm:"column:id;primary_key"`
	//Username 用户名，创建自定义验证器： 长度在 6-20 之间，且不能重复，只能包含大小写字母，数字，下划线；第一个字符必须是字母；
	Username string `json:"username" gorm:"column:username;unique" binding:"usernameValid"`
	//Password 密码， 长度在 6-20 之间，只能包含字母，数字，下划线；
	Password string `json:"password" gorm:"column:password" binding:"passwordValid"`
	//Email 邮箱，不能为空， 必须是邮箱格式，且不能重复；
	Email    string `json:"email" gorm:"column:email;unique" binding:"required,email"`
	Salt     string `json:"salt" gorm:"column:salt"`
	CreateAt int64  `json:"create_at" gorm:"column:create_at"`
	UpdateAt int64  `json:"update_at" gorm:"column:update_at"`
}

func (u *UserModelImpl) TableName() string {
	return "users"
}

func New(attrs ...UserModelAttrFunc) *UserModelImpl {
	u := &UserModelImpl{}
	UserModelAttrFuncs(attrs).apply(u)
	return u
}

func (u *UserModelImpl) Mutate(attrs ...UserModelAttrFunc) *UserModelImpl {
	UserModelAttrFuncs(attrs).apply(u)
	return u
}

// 生成密码
func (u *UserModelImpl) GeneratePassword() error {
	// 使用 bcrypt 生成密码, bcrypt.DefaultCost 表示默认的加密强度，值越大加密强度越大，但是会消耗更多的资源

	pas, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(pas)
	return nil
}
