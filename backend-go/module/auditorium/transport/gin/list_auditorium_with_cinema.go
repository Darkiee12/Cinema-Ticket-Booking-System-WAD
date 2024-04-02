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

// ListAuditoriumWithCinemaID
// @Summary List auditoriums with cinema ID
// @Description List auditoriums with cinema ID
// @Tags auditoriums
// @ID list-auditoriums-with-cinema-id
// @Accept  json
// @Produce  json
// @Param id path string true "Cinema ID"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param cursor query string false "Cursor"
// @Param cinema_name query string false "Cinema Name"
// @Success 200 {object} common.successRes{data=[]auditoriummodel.Auditorium}
// @Router /cinemas/{id}/auditoriums [get]
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

// ListAuditoriumWithCinemaName
// @Summary List auditoriums with cinema name
// @Description List auditoriums with cinema name
// @Tags auditoriums
// @ID list-auditoriums-with-cinema-name
// @Accept  json
// @Produce  json
// @Param name path string true "Cinema Name"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param cursor query string false "Cursor"
// @Success 200 {object} common.successRes{data=[]auditoriummodel.Auditorium}
// @Router /cinemas/{name}/auditoriums [get]
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
