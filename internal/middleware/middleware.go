package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Skipper func(*gin.Context) bool

type Allow struct {
	Prefix string
	Method string
}

var allowList = make([]*Allow, 0)

func AddAllow(path string, method string) {
	allowList = append(allowList, &Allow{
		Prefix: path,
		Method: method,
	})
}

// AllowSkipper 检查请求方法和路径是否包含指定的前缀，如果包含则跳过
func AllowSkipper() Skipper {
	return func(ctx *gin.Context) bool {
		path := ctx.Request.URL.Path
		pathLen := len(path)
		method := strings.ToUpper(ctx.Request.Method)

		for _, p := range allowList {
			if pl := len(p.Prefix); pathLen >= pl && path[:pl] == p.Prefix && strings.ToUpper(p.Method) == method {
				return true
			}
		}
		return false
	}
}
