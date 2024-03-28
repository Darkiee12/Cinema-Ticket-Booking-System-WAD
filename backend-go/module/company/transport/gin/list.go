package gincompany

import (
	"cinema/common"
	"cinema/component/appctx"
	companybusiness "cinema/module/company/biz"
	companymodel "cinema/module/company/model"
	companystore "cinema/module/company/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCompany(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter companymodel.Filter
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

		store := companystore.NewSQLStore(db)
		biz := companybusiness.NewListCompanyBusiness(store)

		result, err := biz.ListCompany(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
