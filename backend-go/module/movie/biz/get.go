package moviebusiness

import (
	moviemodel "cinema/module/movie/model"
	"context"
)

type GetMoviesStore interface {
	GetMovies(
		context context.Context,
		imdbID []string,
	) ([]moviemodel.Movie, error)
}
type getMoviesBusiness struct {
	store GetMoviesStore
}

func NewGetMoviesBusiness(store GetMoviesStore) *getMoviesBusiness {
	return &getMoviesBusiness{store: store}
}
func (business *getMoviesBusiness) GetMovies(
	context context.Context,
	imdbID []string,
) ([]moviemodel.Movie, error) {
	res, err := business.store.GetMovies(context, imdbID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
