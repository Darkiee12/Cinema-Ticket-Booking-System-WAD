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

// CreateCinema gin handler for create cinema
// input:  {name, address, email, phone_number}
func CreateCinema(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		var data cinemamodel.CinemaCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := cinemastore.NewSQLStore(db)
		biz := cinemabuisness.NewCreateCinemaBusiness(store)

		if err := biz.CreateCinema(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeID.String()))
	}
}
