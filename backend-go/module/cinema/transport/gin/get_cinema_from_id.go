package gincinema

import (
	"cinema/common"
	"cinema/component/appctx"
	cinemabuisness "cinema/module/cinema/biz"
	cinemastore "cinema/module/cinema/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCinemaWithID
// @Summary Get a cinema by its ID
// @Description Returns a single cinema
// @Tags cinemas
// @ID get-cinema-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Cinema ID"
// @Success 200 {object} common.successRes{data=cinemamodel.Cinema}
// @Router /cinemas/{id} [get]
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
		biz := cinemabuisness.NewFindCinemaBiz(storage)

		data, err := biz.FindCinemaById(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data.Mask(true)

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(data))
	}
}
