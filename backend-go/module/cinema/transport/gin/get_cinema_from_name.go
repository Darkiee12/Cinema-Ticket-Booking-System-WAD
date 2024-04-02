package gincinema

import (
	"cinema/common"
	"cinema/component/appctx"
	cinemabuisness "cinema/module/cinema/biz"
	cinemastore "cinema/module/cinema/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCinemaWithName
// @Summary Get a cinema by its name
// @Description Returns a single cinema
// @Tags cinemas
// @ID get-cinema-by-name
// @Accept  json
// @Produce  json
// @Param name path string true "Cinema Name"
// @Success 200 {object} common.successRes{data=cinemamodel.Cinema}
// @Router /cinemas/name/{name} [get]
func GetCinemaWithName(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()
		storage := cinemastore.NewSQLStore(db)
		biz := cinemabuisness.NewFindCinemaBiz(storage)

		data, err := biz.FindCinemaByName(c.Request.Context(), c.Param("name"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		data.Mask(true)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
