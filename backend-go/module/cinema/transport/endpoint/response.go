package cinemaendpoint

import (
	"cinema/common"
	cinemamodel "cinema/module/cinema/model"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListCinemaResponse struct {
	Data   []cinemamodel.Cinema `json:"data"`
	Paging common.Paging        `json:",inline"`
	Filter cinemamodel.Filter   `json:",inline"`
}

func EncodeListCinemaResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(ListCinemaResponse)
	c, ok := ctx.Value("ginContext").(*gin.Context)
	if !ok {
		panic("impossible to fail")
	}
	c.JSON(http.StatusOK, common.NewSuccessResponse(resp.Data, resp.Paging, resp.Filter))
	return nil
}
