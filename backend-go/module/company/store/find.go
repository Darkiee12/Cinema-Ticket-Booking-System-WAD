package companystore

import (
	"cinema/common"
	companymodel "cinema/module/company/model"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (store *sqlStore) FindCompany(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*companymodel.Company, error) {
	var data companymodel.Company

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
