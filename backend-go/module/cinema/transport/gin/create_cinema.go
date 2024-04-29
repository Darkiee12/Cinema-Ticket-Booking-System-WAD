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

// CreateCinema
// @Summary Create a cinema
// @Description Create a new cinema
// @Tags cinemas
// @ID create-cinema
// @Accept  json
// @Produce  json
// @Param cinema body cinemamodel.CinemaCreate true "Cinema"
// @Success 200 {object} common.successRes{data=string}
// @Security BearerAuth
// @Router /cinemas [post]
func CreateCinema(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := ctx.GetMainDBConnection()

		var data cinemamodel.CinemaCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := cinemastore.NewSQLStore(db)
		biz := cinemabuisness.NewCreateCinemaBusiness(store)
		data.OwnerID = requester.GetUserId()
		if err := biz.CreateCinema(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(data.FakeID.String()))
	}
}
