package showstore

import (
	"cinema/common"
	showmodel "cinema/module/show/model"
	"context"
)

func (store *sqlStore) GetShows(_ context.Context, id []int) ([]showmodel.Show, error) {
	var result []showmodel.Show

	db := store.db.Table(showmodel.TableName)

	if len(id) > 0 {
		db = db.Where("id in (?)", id)

		if err := db.Find(&result).Error; err != nil {
			return nil, common.ErrDB(err)
		}
		return result, nil
	} else {
		return nil, nil
	}
}
