package auditoriumseatsstore

import (
	"cinema/common"
	auditoriumseatsmodel "cinema/module/auditorium_seat/model"
	"context"
)

func (store *sqlStore) ListSeatsWithCondition(
	_ context.Context,
	filter *auditoriumseatsmodel.Filter,
	moreKeys ...string,
) ([]auditoriumseatsmodel.AuditoriumSeat, error) {
	var result []auditoriumseatsmodel.AuditoriumSeat

	db := store.db.Table(auditoriumseatsmodel.TableName)

	if f := filter; f != nil {
		if len(f.Type) > 0 {
			db = db.Where("type in (?)", f.Type)
		}
		if f.AuditoriumID > 0 {
			db = db.Where("auditorium_id = ?", f.AuditoriumID)
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
