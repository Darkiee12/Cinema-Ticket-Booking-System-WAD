package cinemastore

import (
	"cinema/common"
	cinemamodel "cinema/module/cinema/model"
	"context"
)

func (store *sqlStore) Create(context context.Context, data *cinemamodel.CinemaCreate) error {
	//data.PrepareForInsert()
	if err := store.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
