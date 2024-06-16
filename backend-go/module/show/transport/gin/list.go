package ginshow

import (
	"cinema/common"
	"cinema/component/appctx"
	showbusiness "cinema/module/show/biz"
	showmodel "cinema/module/show/model"
	showstore "cinema/module/show/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListShow
// @Summary List shows
// @Description List shows
// @Tags shows
// @ID list-shows
// @Accept  json
// @Produce  json
// @Param imdbID query string false "IMDB ID"
// @Param date query string false "Date"
// @Param startTime query string false "Start time"
// @Success 200 {object} common.successRes{data=[]showmodel.Show}
// @Router /shows [get]
func ListShow(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		var filter showmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := showstore.NewSQLStore(db)
		biz := showbusiness.NewListShowsBusiness(store)

		result, err := biz.ListShows(c.Request.Context(), &filter)
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(result); i++ {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, filter))
	}
}
