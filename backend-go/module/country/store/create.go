package countrystore

import (
	"cinema/common"
	countrymodel "cinema/module/country/model"
	"context"
)

func (store *sqlStore) Create(context context.Context, data *countrymodel.Country) error {
	//data.PrepareForInsert()
	if err := store.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
