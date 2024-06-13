package movierepository

import (
	"cinema/component/transactor"
	moviemodel "cinema/module/movie/model"
	moviesgenresmodel "cinema/module/movies_genres/model"
	"context"
)

type CreateMovieStore interface {
	Create(context.Context, *moviemodel.Movie) error
	transactor.Transactor
}

type CreateMoviesGenresStore interface {
	Create(context.Context, *moviesgenresmodel.MovieGenre) error
	transactor.Transactor
}

type createMovieRepo struct {
	store             CreateMovieStore
	moviesGenresStore CreateMoviesGenresStore
}

func NewCreateMovieRepo(store CreateMovieStore, moviesGenresStore CreateMoviesGenresStore) *createMovieRepo {
	return &createMovieRepo{store: store, moviesGenresStore: moviesGenresStore}
}

func (repo *createMovieRepo) CreateMovie(ctx context.Context, data *moviemodel.Movie) error {
	return repo.store.WithinTransaction(ctx, func(TxCtx context.Context) error {
		if err := repo.store.Create(TxCtx, data); err != nil {
			return err
		}
		if len(data.Genres) > 0 {
			for _, genre := range data.Genres {
				if err := repo.moviesGenresStore.Create(TxCtx, &moviesgenresmodel.MovieGenre{
					ImdbID:  data.ImdbID,
					GenreID: genre.Id,
				}); err != nil {
					return err
				}
			}
		}
		return nil
	})
}
