package ginshow

import (
	"cinema/common"
	"cinema/component/appctx"
	showbusiness "cinema/module/show/biz"
	showstore "cinema/module/show/store"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetShowWithID
// @Summary Get a show with ID
// @Description Get a show with ID
// @Tags shows
// @ID get-show-with-id
// @Accept  json
// @Produce  json
// @Param id path int true "Show ID"
// @Success 200 {object} common.successRes{data=showmodel.Show}
// @Router /shows/{id} [get]
func GetShowWithID(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("restaurant_id"))

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}

		db := ctx.GetMainDBConnection()
		storage := showstore.NewSQLStore(db)
		biz := showbusiness.NewFindShowBiz(storage)

		data, err := biz.FindShowById(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(data))
	}
}
