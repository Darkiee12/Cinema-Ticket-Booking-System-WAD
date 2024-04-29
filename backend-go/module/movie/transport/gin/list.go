package ginmovie

import (
	"cinema/common"
	"cinema/component/appctx"
	moviebusiness "cinema/module/movie/biz"
	moviemodel "cinema/module/movie/model"
	moviestore "cinema/module/movie/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListMovie
// @Summary List movies
// @Description List movies
// @Tags movies
// @ID list-movies
// @Accept  json
// @Produce  json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param cursor query string false "Cursor"
// @Success 200 {object} common.successRes{data=[]moviemodel.Movie}
// @Router /movies [get]
func ListMovie(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter moviemodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := moviestore.NewSQLStore(db)
		biz := moviebusiness.NewListMovieBusiness(store)

		result, err := biz.ListMovies(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
