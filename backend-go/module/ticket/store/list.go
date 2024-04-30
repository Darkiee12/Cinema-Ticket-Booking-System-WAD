package ticketstore

import (
	"cinema/common"
	ticketmodel "cinema/module/ticket/model"
	"context"
)

func (store *sqlStore) ListTicketsWithCondition(
	_ context.Context,
	filter *ticketmodel.Filter,
	moreKeys ...string,
) ([]ticketmodel.Ticket, error) {
	var result []ticketmodel.Ticket

	db := store.db.Table(ticketmodel.TableName)

	if f := filter; f != nil {
		if f.ShowID > 0 {
			db = db.Where("show_id = ?", f.ShowID)
		}
		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
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
