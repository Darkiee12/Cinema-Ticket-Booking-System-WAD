package ginmovie

import (
	"cinema/common"
	"cinema/component/appctx"
	moviebusiness "cinema/module/movie/biz"
	moviemodel "cinema/module/movie/model"
	movierepository "cinema/module/movie/repository"
	moviestore "cinema/module/movie/store"
	moviesgenresstore "cinema/module/movies_genres/store"
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
		moviesGenresStore := moviesgenresstore.NewSQLStore(db)
		repo := movierepository.NewCreateMovieRepo(store, moviesGenresStore)
		biz := moviebusiness.CreateMovieRepo(repo)

		if err := biz.CreateMovie(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse("true"))
	}
}
