package middleware

import (
	"eos-layout/internal/repository"
	"eos-layout/internal/status"
	"eos-layout/pkg/jwt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const TokenKey = "X-Token"

func Authorization(repo repository.Repository, skipper ...Skipper) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(skipper) > 0 && skipper[0](ctx) {
			ctx.Next()
			return
		}
		// 处理token
		uuid := ""
		if t := ctx.GetHeader(TokenKey); len(t) != 0 {
			token, ok := jwt.ParseToken(t)
			if !ok {
				ctx.JSON(status.ErrorAuthorize.Response())
				ctx.Abort()
				return
			}
			expire, _ := strconv.ParseInt(token["exp"], 10, 64)
			exp := time.Unix(expire, 0)
			ok = exp.After(time.Now())
			if !ok {
				ctx.JSON(status.ErrorAuthorize.Response())
				ctx.Abort()
				return
			}
			uuid = token["uuid"]
			if len(uuid) == 0 {
				ctx.JSON(status.ErrorAuthorize.Response())
				ctx.Abort()
				return
			}
			id, err := repo.Get(ctx, uuid)
			if err != nil {
				ctx.JSON(status.ErrorAuthorize.Response())
				ctx.Abort()
				return
			}
			tid, err := strconv.ParseInt(string(id), 10, 64)
			if err != nil {
				ctx.JSON(status.ErrorAuthorize.Response())
				ctx.Abort()
				return
			}
			ctx.Set("TID", tid)
		}
		ctx.Next()
	}
}
