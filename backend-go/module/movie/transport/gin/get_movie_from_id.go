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
		//id, err := strconv.Atoi(c.Param("restaurant_id"))

		imdbId := c.Param("imdb_id")

		db := ctx.GetMainDBConnection()
		storage := moviestore.NewSQLStore(db)
		biz := moviebusiness.FindMovieStore(storage)

		data, err := biz.FindMovie(c.Request.Context(), map[string]interface{}{"imdb_id": imdbId})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(data))
	}
}
