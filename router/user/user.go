package user

import (
	"github.com/gin-gonic/gin"
	"qh-gin-api/api/v1"
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
