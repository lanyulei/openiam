package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"openiam/pkg/jwtauth"
	"openiam/pkg/tools/respstatus"
	"openiam/server/audit"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/logger"
	"github.com/lanyulei/toolkit/response"
	"github.com/spf13/viper"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			data []byte
		)

		// 获取环境变量 NEBULA_DEBUG
		if os.Getenv("NEBULA_DEBUG") == "true" {
			if c.Request.Method != "GET" && !strings.HasPrefix(c.Request.URL.Path, "/planet/apis/v1") {
				response.Error(c, nil, response.Response{
					Code:    40000,
					Message: "演示环境不允许此操作",
				})
				c.Abort()
				return
			}
		}

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Error(c, nil, respstatus.AuthorizationNullError)
			c.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Error(c, nil, respstatus.AuthorizationFormatError)
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwtauth.ParseToken(parts[1], viper.GetString("jwt.http.secret"))
		if err != nil {
			response.Error(c, err, respstatus.InvalidTokenError)
			c.Abort()
			return
		}

		if c.Request.Method != RequestMethodGet && c.Request.Method != RequestMethodOptions {
			// body
			data, err = c.GetRawData()
			if err != nil {
				logger.Errorf("failed to read request body, error: %v", err.Error())
				return
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(data))

			if isJSONValid(data) {
				go func(c *gin.Context, username string, data []byte) {
					err = audit.Create(c, username, data)
					if err != nil {
						logger.Error(err)
					}
				}(c, mc.Username, data)
			} else {
				logger.Errorf("failed to write audit record, request body is not valid json data")
			}
		}

		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Set("userId", mc.UserId)
		c.Set("isAdmin", mc.IsAdmin)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

func isJSONValid(data []byte) bool {
	var raw json.RawMessage

	// 尝试将数据解析为 json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return false
	}

	return true
}
