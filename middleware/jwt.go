package middleware

import (
	"errors"
	"github.com/lmb1113/qh-gin-api/global"
	"github.com/lmb1113/qh-gin-api/model/common/response"
	"github.com/lmb1113/qh-gin-api/pkg/jwt"
	"github.com/lmb1113/qh-gin-api/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	jwt2 "github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := jwt.GetToken(c)
		if token == "" {
			response.NoAuth("未登录或非法访问", c)
			c.Abort()
			return
		}
		j := jwt.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, jwt.TokenExpired) {
				response.NoAuth("授权已过期", c)
				jwt.ClearToken(c)
				c.Abort()
				return
			}
			response.NoAuth(err.Error(), c)
			jwt.ClearToken(c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.QGA_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = jwt2.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			jwt.SetToken(c, newToken, int(dr.Seconds()))
			if global.QGA_CONFIG.System.UseMultipoint {
				// 记录新的活跃jwt
				_ = jwt.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Next()

		if newToken, exists := c.Get("new-token"); exists {
			c.Header("new-token", newToken.(string))
		}
		if newExpiresAt, exists := c.Get("new-expires-at"); exists {
			c.Header("new-expires-at", newExpiresAt.(string))
		}
	}
}
