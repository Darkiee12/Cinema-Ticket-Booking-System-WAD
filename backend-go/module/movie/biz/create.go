package moviebusiness

import (
	"cinema/common"
	cinemamodel "cinema/module/cinema/model"
	moviemodel "cinema/module/movie/model"
	"context"
)

type CreateMovieRepo interface {
	CreateMovie(ctx context.Context, data *moviemodel.Movie) error
}

type createMovieBusiness struct {
	store CreateMovieRepo
}

func NewCreateMovieBusiness(store CreateMovieRepo) *createMovieBusiness {
	return &createMovieBusiness{store: store}
}

func (business *createMovieBusiness) CreateMovie(context context.Context, data *moviemodel.Movie) error {
	if err := business.store.CreateMovie(context, data); err != nil {
		return common.ErrCannotCreateEntity(cinemamodel.EntityName, err)
	}
	return nil
}
