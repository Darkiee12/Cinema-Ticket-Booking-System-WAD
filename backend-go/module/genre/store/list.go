package genrestore

import (
	"cinema/common"
	genremodel "cinema/module/genre/model"
	"context"
)

// ListGenres list all genres
func (store *sqlStore) ListGenres(
	_ context.Context,
	moreKeys ...string,
) ([]genremodel.Genre, error) {
	var result []genremodel.Genre

	db := store.db.Table(genremodel.TableName)
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
