package ginuser

import (
	"cinema/common"
	"cinema/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfile(_ appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)
		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(data))
	}
}
