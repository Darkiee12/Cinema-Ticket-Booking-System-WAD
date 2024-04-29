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

// CreateMovie
// @Summary Create a movie
// @Description Create a movie
// @Tags movies
// @ID create-movie
// @Accept  json
// @Produce  json
// @Param movie body moviemodel.Movie true "Movie"
// @Success 200 {object} common.successRes{data=string}
// @Security ApiKeyAuth
// @Router /movies [post]
func CreateMovie(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		var data moviemodel.Movie

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := moviestore.NewSQLStore(db)
		biz := moviebusiness.CreateMovieStore(store)

		if err := biz.Create(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse("true"))
	}
}
