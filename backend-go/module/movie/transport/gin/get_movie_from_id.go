package ginmovie

import (
	"cinema/common"
	"cinema/component/appctx"
	moviebusiness "cinema/module/movie/biz"
	moviestore "cinema/module/movie/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetMovieWithID
// @Summary Get a movie with ID
// @Description Get a movie with ID
// @Tags movies
// @ID get-movie-with-id
// @Accept  json
// @Produce  json
// @Param imdb_id path string true "IMDB ID"
// @Success 200 {object} common.successRes{data=moviemodel.Movie}
// @Router /movies/{imdb_id} [get]
func GetMovieWithID(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		imdbId := c.Param("imdb_id")

		db := ctx.GetMainDBConnection()
		storage := moviestore.NewSQLStore(db)
		biz := moviebusiness.NewFindMovieBiz(storage)

		data, err := biz.FindMovieById(c.Request.Context(), imdbId)
		if err != nil {
			panic(err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(data))
	}
}
