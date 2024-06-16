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

// GetTicketsByUser
// @Summary Get tickets by user
// @Description Get tickets by user
// @Tags tickets
// @ID get-tickets-by-user
// @Accept  json
// @Produce  json
// @Success 200 {object} common.successRes{data=[]ticketmodel.Ticket}
// @Security ApiKeyAuth
// @Router /tickets/user [get]
func GetTicketsByUser(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		db := ctx.GetMainDBConnection()

		storage := ticketstore.NewSQLStore(db)
		biz := ticketbusiness.NewListTicketsBusiness(storage)
		tickets, err := biz.ListTicketsByUser(c.Request.Context(), &ticketmodel.Filter{UserID: requester.GetUserId()})
		if err != nil {
			return
		}

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(tickets))
	}
}
