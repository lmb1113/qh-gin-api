package middleware

import (
	"github.com/lmb1113/qh-gin-api/global"
	"github.com/lmb1113/qh-gin-api/model/common/response"
	"net"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const stack = true

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		var id = c.Request.Header.Get("X-Request-ID")
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					global.QGA_LOG.Error(c.Request.URL.Path,
						zap.String("id", id),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					global.QGA_LOG.Error("[Recovery from panic]",
						zap.String("id", id),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
					if gin.IsDebugging() {
						debug.PrintStack()
					}
				} else {
					global.QGA_LOG.Error("[Recovery from panic]",
						zap.String("id", id),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				var data = make(map[string]interface{})
				if gin.IsDebugging() {
					data["header"] = c.Request.Header
					data["request"] = string(httpRequest)
					data["stack"] = string(debug.Stack())
				}
				response.FailWithMessage("未知错误:"+id, c)
			}
		}()
		c.Next()
	}
}
