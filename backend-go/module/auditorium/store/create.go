package auditoriumstore

import (
	"cinema/common"
	auditoriummodel "cinema/module/auditorium/model"
	"context"
)

func (store *sqlStore) Create(context context.Context, data *auditoriummodel.AuditoriumCreate) error {
	//data.PrepareForInsert()
	if err := store.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
