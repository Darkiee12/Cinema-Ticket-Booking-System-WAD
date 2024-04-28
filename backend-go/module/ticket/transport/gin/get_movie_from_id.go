package ginmovie

import (
	"cinema/common"
	"cinema/component/appctx"
	moviebusiness "cinema/module/movie/biz"
	moviestore "cinema/module/movie/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMovieWithID(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("restaurant_id"))

		imdb_id := c.Param("imdb_id")

		db := ctx.GetMainDBConnection()
		storage := moviestore.NewSQLStore(db)
		biz := moviebusiness.FindMovieStore(storage)

		data, err := biz.FindMovie(c.Request.Context(), map[string]interface{}{"imdb_id": imdb_id})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
