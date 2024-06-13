package genrestore

import (
	"cinema/common"
	genremodel "cinema/module/genre/model"
	showmodel "cinema/module/show/model"
	"context"
)

func (store *sqlStore) GetGenres(_ context.Context, ids []int) ([]genremodel.Genre, error) {
	var result []genremodel.Genre
	db := store.db.Table(showmodel.TableName)
	if len(ids) > 0 {
		if err := db.Where("id in (?)", ids).Find(&result).Error; err != nil {
			return nil, common.ErrDB(err)
		}
		return result, nil
	} else {
		return nil, nil
	}
}
