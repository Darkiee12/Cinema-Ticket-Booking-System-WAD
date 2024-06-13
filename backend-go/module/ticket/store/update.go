package ticketstore

import (
	"cinema/common"
	ticketmodel "cinema/module/ticket/model"
	"context"
)

func (store *sqlStore) UpdateTickets(
	_ context.Context,
	cond map[string]interface{},
	data *ticketmodel.TicketUpdate,
) error {
	if err := store.db.Where(cond).
		Select("Status", "UserID").
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
