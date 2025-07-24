package user

import (
	"github.com/gin-gonic/gin"
	"qh-gin-api/model/user/response"
)

type UsersApi struct {
}

// Login
// @Tags      User
// @Summary   用户登录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body    request.LoginRequest  true "手机号 邮箱 密码"
// @Success   200   {object} response.LoginResponse  "登录结果"
// @Router    /user/login [post]
func (u *UsersApi) Login(c *gin.Context) {
	_ = response.LoginResponse{}
	return
}

// Register
// @Tags      User
// @Summary   用户注册
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      nil                                           true  "ID, 用户名, 昵称, 头像链接"
// @Success   200   {object}  nil  "用户信息"
// @Router    /user/register [post]
func (u *UsersApi) Register(c *gin.Context) {
	return
}
