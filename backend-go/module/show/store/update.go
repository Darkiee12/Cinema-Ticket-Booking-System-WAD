package showstore

import (
	"cinema/common"
	showmodel "cinema/module/show/model"
	"context"
)

func (store *sqlStore) UpdateShow(
	ctx context.Context,
	cond map[string]interface{},
	data *showmodel.ShowUpdate,
) error {
	if err := store.db.Where(cond).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
