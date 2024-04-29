package gincinema

import (
	"cinema/common"
	"cinema/component/appctx"
	cinemabuisness "cinema/module/cinema/biz"
	cinemamodel "cinema/module/cinema/model"
	cinemastore "cinema/module/cinema/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListCinema
// @Summary List all cinemas
// @Description Returns a list of cinemas
// @Tags cinemas
// @ID list-cinemas
// @Accept  json
// @Produce  json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param cursor query string false "Cursor"
// @Param owner_id query string false "Owner ID"
// @Success 200 {object} common.successRes{data=[]cinemamodel.Cinema}
// @Router /cinemas [get]
//
//	func ListCinema(ctx appctx.AppContext) gin.HandlerFunc {
//		return func(c *gin.Context) {
//			httptransport.NewServer(
//				cinemaendpoint.NewListCinemaEndpoint(ctx),
//				cinemaendpoint.DecodeListCinemaRequest,
//				cinemaendpoint.EncodeListCinemaResponse,
//			).ServeHTTP(c.Writer, c.Request.WithContext(context.WithValue(c, "ginContext", c)))
//		}
//	}
func ListCinema(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter cinemamodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if fakeID, ok := c.GetQuery("owner_id"); ok {
			uid, err := common.FromBase58(fakeID)
			if err != nil {
				panic(common.ErrInvalidRequest(err))
			}
			filter.OwnerID = int(uid.GetLocalID())
		}

		filter.Status = []int{1}

		store := cinemastore.NewSQLStore(db)
		biz := cinemabuisness.NewListCinemaBusiness(store)

		result, err := biz.ListCinema(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
