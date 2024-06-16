package showstore

import (
	"cinema/common"
	showmodel "cinema/module/show/model"
	"context"
)

func (store *sqlStore) Count(_ context.Context, condition map[string]interface{}) (int64, error) {
	var count int64
	db := store.db.Table(showmodel.TableName)

	for key, value := range condition {
		switch key {
		case "date":
			db = db.Where("date = ?", value)
		case "start_time":
			db = db.Where("start_time >= ?", value)
		case "auditorium_id":
			db = db.Where("auditorium_id = ?", value)
		case "end_time":
			db = db.Where("end_time <= ?", value)
		}
	}

	if err := db.Count(&count).Error; err != nil {
		return 0, common.ErrDB(err)
	}
	return count, nil
}
