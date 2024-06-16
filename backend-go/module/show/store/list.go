package showstore

import (
	"cinema/common"
	showmodel "cinema/module/show/model"
	"context"
	"strings"
)

func (store *sqlStore) ListShowsWithCondition(
	_ context.Context,
	filter *showmodel.Filter,
	moreKeys ...string,
) ([]showmodel.Show, error) {
	var result []showmodel.Show
	db := store.db.Table(showmodel.TableName)
	if f := filter; f != nil {
		f.ImdbID = strings.TrimSpace(f.ImdbID)

		if f.Date != nil {
			db = db.Where("date = ?", f.Date)
			if f.StartTime != nil {
				db = db.Where("start_time >= ?", f.StartTime)
			}
		}
		if f.ImdbID != "" {
			db = db.Where("imdb_id = ?", f.ImdbID)
		}
		if f.Date == nil && f.ImdbID == "" {
			date := common.NowDate()
			db = db.Where("date = ?", (&date).String())
			if f.StartTime != nil {
				db = db.Where("start_time >= ?", f.StartTime)
			}
		}
	} else {
		date := common.NowDate()
		db = db.Where("date = ?", (&date).String())
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Order("start_time ASC").
		Find(&result).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
