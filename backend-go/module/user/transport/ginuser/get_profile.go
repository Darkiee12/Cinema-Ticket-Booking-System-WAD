package ginuser

import (
	"cinema/common"
	"cinema/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetProfile
// @Summary Get profile
// @Description Get profile
// @Tags users
// @ID get-profile
// @Accept  json
// @Produce  json
// @Success 200 {object} common.successRes{data=usermodel.User}
// @Security ApiKeyAuth
// @Router /profile [get]
func GetProfile(_ appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)
		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(data))
	}
}
