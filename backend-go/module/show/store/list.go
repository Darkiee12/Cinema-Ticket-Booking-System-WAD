package showstore

import (
	"cinema/common"
	cinemamodel "cinema/module/cinema/model"
	showmodel "cinema/module/show/model"
	"context"
	"time"
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
			db = db.Where("start_time BETWEEN ? AND ?", f.StartTime, f.StartTime.Add(time.Hour*12))
		}
		if f.Date != nil {
			db = db.Where("date BETWEEN ? AND ?", f.Date, f.Date.Add(time.Hour*24))
		}
		if f.ImdbID != "" {
			db = db.Where("imdb_id = ?", f.ImdbID)
		}
		if f.StartTime == nil && f.Date == nil && f.ImdbID != "" {
			db = db.Where("date BETWEEN ? AND ?", time.Now(), time.Now().Add(time.Hour*24))
		}
	} else {
		db = db.Where("date BETWEEN ? AND ?", time.Now(), time.Now().Add(time.Hour*24))
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Find(&result).
		Order("start_time ASC").
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
