package cinemaendpoint

import (
	"cinema/component/appctx"
	cinemabuisness "cinema/module/cinema/biz"
	cinemastore "cinema/module/cinema/store"
	"context"
	"github.com/go-kit/kit/endpoint"
)

func ListCinemaEndpoint(appCtx appctx.AppContext) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		db := appCtx.GetMainDBConnection()

		req := request.(ListCinemaRequest)

		pagingData := req.Paging

		pagingData.Fulfill()

		filter := req.Filter

		filter.Status = []int{1}

		store := cinemastore.NewSQLStore(db)
		biz := cinemabuisness.NewListCinemaBusiness(store)

		result, err := biz.ListCinema(ctx, &filter, &pagingData)
		if err != nil {
			return nil, err
		}

		for i := range result {
			result[i].Mask(false)
		}

		return ListCinemaResponse{
			Data:   result,
			Paging: pagingData,
			Filter: filter,
		}, nil
	}
}
