package ginmoviesgenres

import (
	"cinema/common"
	"cinema/component/appctx"
	moviesgenresbusiness "cinema/module/movies_genres/biz"
	moviesgenresmodel "cinema/module/movies_genres/model"
	moviesgenresstore "cinema/module/movies_genres/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListMoviesByGenres
// @Summary List movies by genres
// @Description List movies by genres
// @Tags movies_genres
// @ID list-movies-by-genres
// @Accept  json
// @Produce  json
// @Param genre_id path string true "Genre ID"
// @Success 200 {object} common.successRes{data=[]moviesgenresmodel.MovieGenre}
// @Router /genres/:genre_id/movies [get]
func ListMoviesByGenres(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()
		var filter moviesgenresmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := moviesgenresstore.NewSQLStore(db)
		biz := moviesgenresbusiness.NewListMoviesGenresBusiness(store)

		result, err := biz.ListMoviesByGenres(c.Request.Context(), &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, filter))
	}
}
