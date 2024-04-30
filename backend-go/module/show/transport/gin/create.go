package ginshow

import (
	"cinema/common"
	"cinema/component/appctx"
	auditoriumbusiness "cinema/module/auditorium/biz"
	auditoriumstore "cinema/module/auditorium/store"
	auditoriumseatsstore "cinema/module/auditorium_seat/store"
	showbusiness "cinema/module/show/biz"
	showmodel "cinema/module/show/model"
	showrepository "cinema/module/show/repository"
	showstore "cinema/module/show/store"
	ticketstore "cinema/module/ticket/store"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateShow
// @Summary Create a show
// @Description Create a show
// @Tags shows
// @ID create-show
// @Accept  json
// @Produce  json
// @Param show body showmodel.Show true "Show"
// @Success 200 {object} common.successRes{data=string}
// @Security ApiKeyAuth
// @Router /shows [post]
func CreateShow(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		db := ctx.GetMainDBConnection()

		var data showmodel.Show

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		uid, err := common.FromBase58(data.AuditoriumFakeID)

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		// Check if the user has permission to create show
		if requester.GetRole() != "admin" {
			store := auditoriumstore.NewSQLStore(db)
			biz := auditoriumbusiness.NewFindAuditoriumBiz(store)
			auditorium, err := biz.FindAuditoriumById(c.Request.Context(), int(uid.GetLocalID()))
			if err != nil {
				return
			}
			if auditorium.Cinema.OwnerID != requester.GetUserId() {
				panic(common.ErrNoPermission(errors.New("user does not have permission to create show")))
			}
		}

		data.AuditoriumID = int64(uid.GetLocalID())

		store := showstore.NewSQLStore(db)
		ticketStore := ticketstore.NewSQLStore(db)
		seatsStore := auditoriumseatsstore.NewSQLStore(db)

		repo := showrepository.NewCreateShowRepo(store, ticketStore, seatsStore)
		biz := showbusiness.NewCreateShowBusiness(repo)

		if err := biz.CreateShow(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse("true"))
	}
}
