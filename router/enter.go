package router

import (
	"qh-gin-api/router/user"
)

type RouterGroup struct {
	User user.UsersRouter
}

var RouterGroupApp = new(RouterGroup)
