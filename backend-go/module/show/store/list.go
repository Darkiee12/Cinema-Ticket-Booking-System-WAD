package showstore

import (
	"cinema/common"
	showmodel "cinema/module/show/model"
	"context"
	"log"
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
		}
		if f.ImdbID != "" {
			db = db.Where("imdb_id = ?", f.ImdbID)
		}
		if f.Date == nil && f.ImdbID == "" {
			date := common.NowDate()
			db = db.Where("date = ?", (&date).String())
		}
	} else {
		date := common.NowDate()
		log.Println("Test: ", (&date).String())
		db = db.Where("date = ?", (&date).String())
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
