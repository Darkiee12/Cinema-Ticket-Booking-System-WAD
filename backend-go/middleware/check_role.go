package middleware

import (
	"cinema/common"
	"cinema/component/appctx"
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckRole(appCtx appctx.AppContext, allowRoles ...string) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		for _, role := range allowRoles {
			if requester.GetRole() == role {
				c.Next()
				return
			}
		}
		panic(common.ErrNoPermission(errors.New("invalid role")))
	}

}
