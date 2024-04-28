package showstore

import (
	"cinema/common"
	cinemamodel "cinema/module/cinema/model"
	showmodel "cinema/module/show/model"
	"context"
)

func (store *sqlStore) ListShowsWithCondition(
	_ context.Context,
	filter *showmodel.Filter,
	moreKeys ...string,
) ([]showmodel.Show, error) {
	var result []showmodel.Show

	db := store.db.Table(cinemamodel.TableName)

	if f := filter; f != nil {
		if f.StartTime != nil {
			db = db.Where("start_time >= ?", f.StartTime)
		}
		if f.Date != nil {
			db = db.Where("date = ?", f.Date)
		}
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Find(&result).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
