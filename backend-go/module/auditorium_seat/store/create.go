package auditoriumseatsstore

import (
	"cinema/common"
	auditoriumseatsmodel "cinema/module/auditorium_seat/model"
	"context"
)

func (store *sqlStore) Create(ctx context.Context, data *auditoriumseatsmodel.AuditoriumSeat) error {
	//data.PrepareForInsert()

	if err := store.model(ctx).Omit("UpdatedAt").Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (store *sqlStore) CreateMultiple(ctx context.Context, data []*auditoriumseatsmodel.AuditoriumSeat) error {
	//data.PrepareForInsert()
	if err := store.model(ctx).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
