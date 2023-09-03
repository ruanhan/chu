package handlers

import (
	"com.github.goscaffold/pkg/data/Getter"
	"com.github.goscaffold/pkg/models/UserModel"
	"com.github.goscaffold/pkg/result"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// Build方法
func (this *UserHandler) Build(r *gin.Engine) {
	r.GET("/user", UserList)
	r.GET("/user/:id", UserDetail)
	r.POST("/user", UserSave)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
}

func UserList(c *gin.Context) {
	//user := UserModel.New()
	//// 验证参数
	//result.Result(c.ShouldBind(user)).Unwrap()
	//
	////OK(c)("get userlist success", "100001", result.Result(test.GetInfo(user.UserId)).Unwrap())
	//
	//if user.UserId > 10 {
	//	ResultWrapper(c)("get userlist success", "100001", "userlist")(OK)
	//} else {
	//	ResultWrapper(c)("get userlist error", "100002", "userlist")(Error)
	//}

	ResultWrapper(c)("get userlist success", "100001", Getter.UserGetter.GetUserList())(OK)
}

func UserDetail(c *gin.Context) {
	id := &struct {
		Id int `uri:"id" binding:"required"`
	}{}
	result.Result(c.ShouldBindUri(id)).Unwrap()
	ResultWrapper(c)("get userdetail success", "100001", Getter.UserGetter.GetUserDetail(id.Id).Unwrap())(OK)
}

func UserSave(c *gin.Context) {
	u := UserModel.New()
	result.Result(c.ShouldBindJSON(u)).Unwrap()
	ResultWrapper(c)("save user", "10086", "true")(OK)
}
