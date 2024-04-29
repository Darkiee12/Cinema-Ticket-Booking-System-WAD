package moviebusiness

import (
	"cinema/common"
	moviemodel "cinema/module/movie/model"
	"context"
)

type ListMovieStore interface {
	ListMoviesWithCondition(
		context context.Context,
		filter *moviemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]moviemodel.Movie, error)
}
type listMovieBusiness struct {
	store ListMovieStore
}

func NewListMovieBusiness(store ListMovieStore) *listMovieBusiness {
	return &listMovieBusiness{store: store}
}
func (business *listMovieBusiness) ListMovies(
	context context.Context,
	filter *moviemodel.Filter,
	paging *common.Paging,
) ([]moviemodel.Movie, error) {
	res, err := business.store.ListMoviesWithCondition(context, filter, paging)
	if err != nil {
		return nil, err
	}
	return res, nil
}
