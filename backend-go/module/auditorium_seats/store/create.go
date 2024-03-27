package auditoriumseatsstore

import (
	"cinema/common"
	auditoriumseatsmodel "cinema/module/auditorium_seats/model"
	"context"
)

func (store *sqlStore) Create(context context.Context, data *auditoriumseatsmodel.AuditoriumSeat) error {
	//data.PrepareForInsert()
	if err := store.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (store *sqlStore) CreateMultiple(context context.Context, data []*auditoriumseatsmodel.AuditoriumSeat) error {
	//data.PrepareForInsert()
	if err := store.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
