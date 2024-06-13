package moviesgenresstore

import (
	"cinema/common"
	auditoriumseatsmodel "cinema/module/auditorium_seat/model"
	moviesgenresmodel "cinema/module/movies_genres/model"
	"context"
)

func (store *sqlStore) ListMoviesGenres(
	_ context.Context,
	filter *moviesgenresmodel.Filter,
	moreKeys ...string,
) ([]moviesgenresmodel.MovieGenre, error) {
	var result []moviesgenresmodel.MovieGenre

	db := store.db.Table(auditoriumseatsmodel.TableName)

	if f := filter; f != nil {
		if len(f.ImdbID) > 0 {
			db = db.Where("imdb_id = ?", f.ImdbID)
		}
		if f.GenreID > 0 {
			db = db.Where("genre_id = ?", f.GenreID)
		}
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Find(&result).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
