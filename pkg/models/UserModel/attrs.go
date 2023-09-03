package UserModel

type UserModelAttrFunc func(u *UserModelImpl)

type UserModelAttrFuncs []UserModelAttrFunc

func WithUserId(id int) UserModelAttrFunc {
	return func(u *UserModelImpl) {
		u.Id = id
	}
}

func WithUserName(name string) UserModelAttrFunc {
	return func(u *UserModelImpl) {
		u.Username = name
	}
}

func WithPassword(pwd string) UserModelAttrFunc {
	return func(u *UserModelImpl) {
		u.Password = pwd
	}
}

func (this UserModelAttrFuncs) apply(u *UserModelImpl) {
	for _, f := range this {
		f(u)
	}
}
