package moviestore

import (
	"cinema/common"
	moviemodel "cinema/module/movie/model"
	"context"
)

func (store *sqlStore) GetMovies(_ context.Context, imdbID []string) ([]moviemodel.Movie, error) {
	var result []moviemodel.Movie

	db := store.db.Table(moviemodel.TableName)

	if len(imdbID) > 0 {
		db = db.Where("imdb_id in (?)", imdbID)
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
