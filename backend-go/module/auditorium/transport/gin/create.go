package ginauditorium

import (
	"cinema/common"
	"cinema/component/appctx"
	auditoriumbusiness "cinema/module/auditorium/biz"
	auditoriummodel "cinema/module/auditorium/model"
	auditoriumrepository "cinema/module/auditorium/repository"
	auditoriumstore "cinema/module/auditorium/store"
	auditoriumseatsstore "cinema/module/auditorium_seats/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateAuditorium input: {name, seats, cinema_id}
// example: curl -X POST http://localhost:8080/v1/auditoriums -d '{"name":"A1","seats":100,"cinema_id":1}'
func CreateAuditorium(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()
		var data auditoriummodel.AuditoriumCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := auditoriumstore.NewSQLStore(db)
		seatsStore := auditoriumseatsstore.NewSQLStore(db)
		repo := auditoriumrepository.NewCreateAuditoriumRepo(store, seatsStore)
		biz := auditoriumbusiness.NewCreateAuditoriumBusiness(repo)

		if err := biz.CreateAuditorium(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeID.String()))
	}
}
