package showstore

import (
	"cinema/common"
	showmodel "cinema/module/show/model"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (store *sqlStore) FindShow(
	_ context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*showmodel.Show, error) {
	var data showmodel.Show

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
