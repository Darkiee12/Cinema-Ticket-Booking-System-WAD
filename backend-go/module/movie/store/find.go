package moviestore

import (
	"cinema/common"
	moviemodel "cinema/module/movie/model"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (store *sqlStore) FindMovie(
	_ context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*moviemodel.Movie, error) {
	var data moviemodel.Movie
	db := store.db.Table(moviemodel.TableName)
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	if err := db.Where(condition).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
