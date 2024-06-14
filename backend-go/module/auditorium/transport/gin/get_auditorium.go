package ginauditorium

import (
	"cinema/common"
	"cinema/component/appctx"
	auditoriumbusiness "cinema/module/auditorium/biz"
	auditoriumstore "cinema/module/auditorium/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAuditoriumWithID
// @Summary Get a auditorium by its ID
// @Description Returns a single auditorium
// @Tags auditoriums
// @ID get-auditorium-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Auditorium ID"
// @Success 200 {object} common.successRes{data=auditoriummodel.Auditorium}
// @Router /auditoriums/{id} [get]
func GetAuditoriumWithID(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		db := ctx.GetMainDBConnection()
		storage := auditoriumstore.NewSQLStore(db)
		biz := auditoriumbusiness.NewFindAuditoriumBiz(storage)

		data, err := biz.FindAuditoriumById(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data.Mask(true)

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(data))
	}
}
