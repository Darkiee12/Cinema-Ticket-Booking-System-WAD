package companystore

import (
	"cinema/common"
	companymodel "cinema/module/company/model"
	"context"
)

func (store *sqlStore) Create(context context.Context, data *companymodel.CompanyCreate) error {
	//data.PrepareForInsert()
	if err := store.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
