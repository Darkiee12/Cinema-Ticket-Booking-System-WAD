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

// UpdateTicket
// @Summary Let user buy a ticket
// @Description Let user buy a ticket
// @Tags tickets
// @ID update-ticket
// @Accept  json
// @Produce  json
// @Param ticket body ticketmodel.TicketUpdate true "Ticket"
// @Security ApiKeyAuth
// @Success 200 {object} common.successRes{data=string}
// @Router /tickets [put]
func UpdateTicket(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()
		var data ticketmodel.TicketUpdate
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		data.UserID = int64(requester.GetUserId())

		store := ticketstore.NewSQLStore(db)
		biz := ticketbusiness.NewUpdateTicketBiz(store)

		if err := biz.SellTicketToCustomer(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(true))
	}
}
