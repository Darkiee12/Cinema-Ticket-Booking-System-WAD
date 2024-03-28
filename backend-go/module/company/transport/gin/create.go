package gincompany

import (
	"cinema/common"
	"cinema/component/appctx"
	companybuisness "cinema/module/company/biz"
	companymodel "cinema/module/company/model"
	companystore "cinema/module/company/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateCompany gin handler for create cinema
func CreateCompany(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		var data companymodel.CompanyCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := companystore.NewSQLStore(db)
		biz := companybuisness.NewCreateCompanyBusiness(store)

		if err := biz.CreateCompany(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeID.String()))
	}
}
