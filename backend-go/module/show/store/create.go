package showstore

import (
	"cinema/common"
	showmodel "cinema/module/show/model"
	"context"
)

func (store *sqlStore) Create(context context.Context, data *showmodel.ShowCreate) error {
	//data.PrepareForInsert()
	if err := store.model(context).Omit("ID").Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
