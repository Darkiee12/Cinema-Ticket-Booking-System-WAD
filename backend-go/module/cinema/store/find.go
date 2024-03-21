package cinemastore

import (
	"cinema/common"
	cinemamodel "cinema/module/cinema/model"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (store *sqlStore) FindCinema(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*cinemamodel.Cinema, error) {
	var data cinemamodel.Cinema

	db := store.db
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
