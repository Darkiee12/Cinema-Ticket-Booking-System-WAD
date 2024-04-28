package ticketstore

import (
	"cinema/common"
	ticketmodel "cinema/module/ticket/model"
	"context"
)

func (store *sqlStore) Create(context context.Context, data *ticketmodel.TicketCreate) error {
	//data.PrepareForInsert()
	if err := store.model(context).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
