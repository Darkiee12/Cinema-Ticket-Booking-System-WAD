package auditoriumstore

import (
	"cinema/common"
	auditoriummodel "cinema/module/auditorium/model"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (store *sqlStore) FindAuditorium(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*auditoriummodel.Auditorium, error) {
	var data auditoriummodel.Auditorium

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
