package ginuser

import (
	"cinema/common"
	"cinema/component/appctx"
	userbiz "cinema/module/user/business"
	"cinema/module/user/usermodel"
	"cinema/module/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateUser
// @Summary Update user profile
// @Description Update user profile
// @Tags users
// @Accept  json
// @Produce  json
// @Param data body usermodel.UserUpdate true "User data"
// @Success 200 {object} common.successRes{data=bool}
// @Security ApiKeyAuth
// @Router /profile [put]
func UpdateUser(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserUpdate
		user := c.MustGet(common.CurrentUser).(common.Requester)

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := ctx.GetMainDBConnection()
		store := userstore.NewSQLStore(db)
		biz := userbiz.NewUpdateUserBiz(store)
		if err := biz.UpdateProfileById(c.Request.Context(), user.GetUserId(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(true))
	}
}
