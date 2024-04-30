package ginuser

import (
	"cinema/common"
	"cinema/component/appctx"
	"cinema/component/hasher"
	"cinema/component/tokenprovider/jwt"
	userbiz "cinema/module/user/business"
	"cinema/module/user/usermodel"
	"cinema/module/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login
// @Summary Login
// @Description Login
// @Tags users
// @ID login
// @Accept  json
// @Produce  json
// @Param cinema body usermodel.UserLogin true "User"
// @Success 200 {object} common.successRes{data=tokenprovider.Token}
// @Router /login [post]
func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey()) //appctx.SecretKey()

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewLoginBusiness(appCtx, store, 60*60*24*30, tokenProvider, md5)
		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(account))
	}
}
