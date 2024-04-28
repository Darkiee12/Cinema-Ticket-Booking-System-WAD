package ticketstore

import (
	"cinema/common"
	ticketmodel "cinema/module/ticket/model"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (store *sqlStore) FindTicket(
	_ context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*ticketmodel.Ticket, error) {
	var data ticketmodel.Ticket

	db := store.db
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	if err := db.Where(condition).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
