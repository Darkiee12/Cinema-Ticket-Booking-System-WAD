package cinemaendpoint

import (
	"cinema/common"
	cinemamodel "cinema/module/cinema/model"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListCinemaRequest struct {
	Paging common.Paging      `json:",inline"`
	Filter cinemamodel.Filter `json:",inline"`
}

func DecodeListCinemaRequest(ctx context.Context, r *http.Request) (interface{}, error) {

	c, ok := ctx.Value("ginContext").(*gin.Context)
	if !ok {
		panic("impossible to fail")
	}

	var pagingData common.Paging
	if err := c.ShouldBind(&pagingData); err != nil {
		panic(common.ErrInvalidRequest(err))
	}

	pagingData.Fulfill()

	var filter cinemamodel.Filter
	if err := c.ShouldBind(&filter); err != nil {
		panic(common.ErrInvalidRequest(err))
	}
	filter.Status = []int{1}

	return ListCinemaRequest{
		Paging: pagingData,
		Filter: filter,
	}, nil
}
