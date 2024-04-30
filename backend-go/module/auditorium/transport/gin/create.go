package ginauditorium

import (
	"cinema/common"
	"cinema/component/appctx"
	auditoriumbusiness "cinema/module/auditorium/biz"
	auditoriummodel "cinema/module/auditorium/model"
	auditoriumrepository "cinema/module/auditorium/repository"
	auditoriumstore "cinema/module/auditorium/store"
	auditoriumseatsstore "cinema/module/auditorium_seat/store"
	cinemabuisness "cinema/module/cinema/biz"
	cinemastore "cinema/module/cinema/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateAuditorium
// @Summary Create an auditorium
// @Description Create an auditorium
// @Tags auditoriums
// @ID create-auditorium
// @Accept  json
// @Produce  json
// @Param auditorium body auditoriummodel.AuditoriumCreate true "Auditorium"
// @Success 200 {object} common.successRes{data=string}
// @Security ApiKeyAuth
// @Router /auditoriums [post]
func CreateAuditorium(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()
		var data auditoriummodel.AuditoriumCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		if requester.GetRole() != "admin" {
			store := cinemastore.NewSQLStore(db)
			biz := cinemabuisness.NewFindCinemaBiz(store)
			cinema, err := biz.FindCinemaByOwnerID(c.Request.Context(), requester.GetUserId())
			if err != nil {
				panic(err)
			}
			(&data).CinemaID = cinema.ID
		}

		store := auditoriumstore.NewSQLStore(db)
		seatsStore := auditoriumseatsstore.NewSQLStore(db)
		repo := auditoriumrepository.NewCreateAuditoriumRepo(store, seatsStore)
		biz := auditoriumbusiness.NewCreateAuditoriumBusiness(repo)

		if err := biz.CreateAuditorium(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(data.FakeID.String()))
	}
}
