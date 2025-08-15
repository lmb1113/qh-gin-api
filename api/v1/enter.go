package v1

import "github.com/lmb1113/qh-gin-api/api/v1/user"

type ApiGroup struct {
	UserApiGroup user.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
