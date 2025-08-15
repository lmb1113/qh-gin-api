package user

import (
	"github.com/lmb1113/qh-gin-api/api/v1"

	"github.com/gin-gonic/gin"
)

type UsersRouter struct {
}

func (u *UsersRouter) InitUsersRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("users")
	var userApi = v1.ApiGroupApp.UserApiGroup.UsersApi
	{
		userRouter.POST("login", userApi.Login)
		userRouter.POST("register", userApi.Register)

	}
}
