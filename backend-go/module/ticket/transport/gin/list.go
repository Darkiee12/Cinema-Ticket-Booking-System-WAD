package ginticket

import (
	"cinema/common"
	"cinema/component/appctx"
	ticketbusiness "cinema/module/ticket/biz"
	ticketmodel "cinema/module/ticket/model"
	ticketstore "cinema/module/ticket/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListTickets
// @Summary List tickets
// @Description List tickets
// @Tags tickets
// @ID list-tickets
// @Accept  json
// @Produce  json
// @Param showID query string false "Show ID"
// @Success 200 {object} common.successRes{data=[]ticketmodel.Ticket}
// @Security ApiKeyAuth
// @Router /tickets [get]
func ListTickets(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		var filter ticketmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter.Validate()

		store := ticketstore.NewSQLStore(db)
		biz := ticketbusiness.NewListTicketsBusiness(store)

		result, err := biz.ListTickets(c.Request.Context(), &filter)
		var pagingData common.Paging
		pagingData.Total = int64(len(result))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, &pagingData, &filter))
	}
}
