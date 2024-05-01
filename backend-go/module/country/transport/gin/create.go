package gincountry

import (
	"cinema/common"
	"cinema/component/appctx"
	countrybusiness "cinema/module/country/biz"
	countrymodel "cinema/module/country/model"
	countrystore "cinema/module/country/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCountry(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		var data countrymodel.Country

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		

		store := countrystore.NewSQLStore(db)
		biz := countrybusiness.NewCreateCountryBusiness(store)

		if err := biz.CreateCountry(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(data.Name))
	}
}
