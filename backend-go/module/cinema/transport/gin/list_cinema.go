package gincinema

import (
	"cinema/common"
	"cinema/component/appctx"
	cinemabuisness "cinema/module/cinema/biz"
	cinemamodel "cinema/module/cinema/model"
	cinemastore "cinema/module/cinema/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCinema(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter cinemamodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if fakeID, ok := c.GetQuery("id"); ok {
			uid, err := common.FromBase58(fakeID)
			if err != nil {
				panic(common.ErrInvalidRequest(err))
			}
			filter.OwnerID = int(uid.GetLocalID())
		}

		filter.Status = []int{1}

		store := cinemastore.NewSQLStore(db)
		biz := cinemabuisness.NewListCinemaBusiness(store)

		result, err := biz.ListCinema(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
