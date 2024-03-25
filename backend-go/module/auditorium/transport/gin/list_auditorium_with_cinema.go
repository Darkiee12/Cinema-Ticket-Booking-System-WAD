package ginauditorium

import (
	"cinema/common"
	"cinema/component/appctx"
	auditoriumbusiness "cinema/module/auditorium/biz"
	auditoriummodel "cinema/module/auditorium/model"
	auditoriumrepository "cinema/module/auditorium/repository"
	auditoriumstore "cinema/module/auditorium/store"
	cinemastore "cinema/module/cinema/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListAuditoriumWithCinemaID(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		db := ctx.GetMainDBConnection()

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter auditoriummodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter.Status = []int{1}
		filter.CinemaID = int(uid.GetLocalID())

		store := auditoriumstore.NewSQLStore(db)
		biz := auditoriumbusiness.NewListAuditoriumBusiness(store)

		result, err := biz.ListAuditorium(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}

func ListAuditoriumWithCinemaName(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		cinemaName := c.Param("name")

		db := ctx.GetMainDBConnection()

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter auditoriummodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter.Status = []int{1}

		store := auditoriumstore.NewSQLStore(db)
		findCinemaStore := cinemastore.NewSQLStore(db)
		repo := auditoriumrepository.NewListAuditoriumWithCinemaNameRepo(store, findCinemaStore)
		biz := auditoriumbusiness.NewListAuditoriumWithCinemaNameBusiness(repo)

		result, err := biz.ListAuditoriumWithCinemaName(c.Request.Context(), &filter, &pagingData, cinemaName)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
