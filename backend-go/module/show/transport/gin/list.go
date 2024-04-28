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

func ListMovie(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		var filter moviemodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := moviestore.NewSQLStore(db)
		biz := moviebusiness.NewListMovieBusiness(store)

		result, err := biz.ListMovies(c.Request.Context(), &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, filter))
	}
}
