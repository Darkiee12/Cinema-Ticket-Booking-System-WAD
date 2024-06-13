package gingenre

import (
	"cinema/common"
	"cinema/component/appctx"
	genrebusiness "cinema/module/genre/biz"
	genrestore "cinema/module/genre/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListGenres list all genres
// @Summary List genres
// @Description List genres
// @Tags genres
// @ID list-genres
// @Accept  json
// @Produce  json
// @Success 200 {object} common.successRes{data=[]genremodel.Genre}
// @Router /genres [get]
func ListGenres(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		store := genrestore.NewSQLStore(db)
		biz := genrebusiness.NewListGenreBusiness(store)

		result, err := biz.ListGenres(c.Request.Context())
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(result))
	}
}
