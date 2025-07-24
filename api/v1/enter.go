package v1

import "qh-gin-api/api/v1/user"

type ApiGroup struct {
	UserApiGroup user.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
