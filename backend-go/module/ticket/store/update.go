package ticketstore

import (
	"cinema/common"
	ticketmodel "cinema/module/ticket/model"
	"context"
)

func (store *sqlStore) UpdateTicket(
	_ context.Context,
	cond map[string]interface{},
	data *ticketmodel.TicketUpdate,
) error {
	if err := store.db.Where(cond).
		Select("SeatNumber", "ShowID", "Status", "UserID").
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
