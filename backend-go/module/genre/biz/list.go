package genrebusiness

import (
	genremodel "cinema/module/genre/model"
	"context"
)

type ListGenreStore interface {
	ListGenres(
		context context.Context,
		moreKeys ...string,
	) ([]genremodel.Genre, error)
}
type listGenreBusiness struct {
	store ListGenreStore
}

func NewListGenreBusiness(store ListGenreStore) *listGenreBusiness {
	return &listGenreBusiness{store: store}
}
func (business *listGenreBusiness) ListGenres(
	context context.Context,
) ([]genremodel.Genre, error) {
	res, err := business.store.ListGenres(context)
	if err != nil {
		return nil, err
	}
	return res, nil
}
