package cinemastore

import (
	"cinema/common"
	cinemamodel "cinema/module/cinema/model"
	"context"
)

func (store *sqlStore) UpdateCinema(
	ctx context.Context,
	cond map[string]interface{},
	data *cinemamodel.Cinema,
) error {
	if err := store.db.Where(cond).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
