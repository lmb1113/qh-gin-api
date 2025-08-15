package router

import (
	"github.com/lmb1113/qh-gin-api/router/user"
)

type RouterGroup struct {
	User user.UsersRouter
}

var RouterGroupApp = new(RouterGroup)
