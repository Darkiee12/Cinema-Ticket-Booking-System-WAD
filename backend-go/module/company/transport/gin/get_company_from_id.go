package gincompany

import (
	"cinema/common"
	"cinema/component/appctx"
	companybusiness "cinema/module/company/biz"
	companystore "cinema/module/company/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCompanyWithID
// @Summary Get a company by its ID
// @Description Returns a single company
// @Tags companies
// @ID get-company
// @Accept  json
// @Produce  json
// @Param id path string true "Company ID"
// @Success 200 {object} common.successRes{data=companymodel.Company}
func GetCompanyWithID(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("restaurant_id"))
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		db := ctx.GetMainDBConnection()
		storage := companystore.NewSQLStore(db)
		biz := companybusiness.NewFindCompanyBiz(storage)

		data, err := biz.FindCompanyById(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data.Mask(true)

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(data))
	}
}
