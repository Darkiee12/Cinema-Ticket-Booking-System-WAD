package auditoriumstore

import (
	"cinema/common"
	auditoriummodel "cinema/module/auditorium/model"
	"context"
)

func (store *sqlStore) Create(ctx context.Context, data *auditoriummodel.AuditoriumCreate) error {
	//data.PrepareForInsert()
	if err := store.model(ctx).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
