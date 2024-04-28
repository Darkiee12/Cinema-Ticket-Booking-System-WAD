package ticketstore

import (
	"cinema/common"
	ticketmodel "cinema/module/ticket/model"
	"context"
)

func (store *sqlStore) UpdateTicket(
	ctx context.Context,
	cond map[string]interface{},
	data *ticketmodel.TicketUpdate,
) error {
	if err := store.db.Where(cond).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
