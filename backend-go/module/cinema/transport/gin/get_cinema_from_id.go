package gincinema

import (
	"cinema/common"
	"cinema/component/appctx"
	cinemabuisness "cinema/module/cinema/biz"
	cinemastore "cinema/module/cinema/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCinemaWithID(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("restaurant_id"))

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		db := ctx.GetMainDBConnection()
		storage := cinemastore.NewSQLStore(db)
		biz := cinemabuisness.NewFindRestaurantBiz(storage)

		data, err := biz.FindCinemaById(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data.Mask(true)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
