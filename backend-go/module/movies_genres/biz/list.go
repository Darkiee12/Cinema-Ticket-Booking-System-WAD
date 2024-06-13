package moviesgenresbusiness

import (
	moviesgenresmodel "cinema/module/movies_genres/model"
	"context"
)

type ListMoviesGenresStore interface {
	ListMoviesGenres(
		_ context.Context,
		filter *moviesgenresmodel.Filter,
		moreKeys ...string,
	) ([]moviesgenresmodel.MovieGenre, error)
}
type listMoviesGenresBusiness struct {
	store ListMoviesGenresStore
}

func NewListMoviesGenresBusiness(store ListMoviesGenresStore) *listMoviesGenresBusiness {
	return &listMoviesGenresBusiness{store: store}
}
func (business *listMoviesGenresBusiness) ListMoviesByGenres(
	context context.Context,
	filter *moviesgenresmodel.Filter,
) ([]moviesgenresmodel.MovieGenre, error) {
	res, err := business.store.ListMoviesGenres(context, filter, "Movies", "Genres")
	if err != nil {
		return nil, err
	}
	return res, nil
}
