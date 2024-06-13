package moviebusiness

import (
	moviemodel "cinema/module/movie/model"
	"context"
)

type FindMovieStore interface {
	FindMovie(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*moviemodel.Movie, error)
}

func NewFindMovieBiz(store FindMovieStore) *findMovieBiz {
	return &findMovieBiz{store: store}
}

type findMovieBiz struct {
	store FindMovieStore
}

func (biz *findMovieBiz) FindMovieById(ctx context.Context, imdbID string) (*moviemodel.Movie, error) {
	data, err := biz.store.FindMovie(ctx, map[string]interface{}{"imdb_id": imdbID}, "Genres")

	if err != nil {
		return nil, err
	}

	return data, nil
}
