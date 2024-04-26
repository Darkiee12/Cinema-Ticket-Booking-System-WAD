package moviebusiness

import (
	"cinema/common"
	cinemamodel "cinema/module/cinema/model"
	moviemodel "cinema/module/movie/model"
	"context"
)

type CreateMovieStore interface {
	Create(context.Context, *moviemodel.Movie) error
}

type createMovieBusiness struct {
	store CreateMovieStore
}

func NewCreateMovieBusiness(store CreateMovieStore) *createMovieBusiness {
	return &createMovieBusiness{store: store}
}

func (business *createMovieBusiness) CreateMovie(context context.Context, data *moviemodel.Movie) error {
	if err := business.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(cinemamodel.EntityName, err)
	}
	return nil
}
