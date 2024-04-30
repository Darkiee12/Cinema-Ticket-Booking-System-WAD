package ginuser

import (
	"cinema/common"
	"cinema/component/appctx"
	"cinema/component/hasher"
	userbiz "cinema/module/user/business"
	"cinema/module/user/usermodel"
	"cinema/module/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register
// @Summary Register
// @Description Register
// @Tags users
// @ID register
// @Accept  json
// @Produce  json
// @Param cinema body usermodel.UserCreate true "User"
// @Success 200 {object} common.successRes{data=string}
// @Router /register [post]
func Register(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		repo := userbiz.NewRegisterBusiness(store, md5)

		if err := repo.Register(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(data.FakeID.String()))
	}
}
