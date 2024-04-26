package moviebusiness

import (
	moviemodel "cinema/module/movie/model"
	"context"
)

type ListMovieStore interface {
	ListCompanyWithCondition(
		context context.Context,
		filter *moviemodel.Filter,
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
) ([]moviemodel.Movie, error) {
	res, err := business.store.ListCompanyWithCondition(context, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}
