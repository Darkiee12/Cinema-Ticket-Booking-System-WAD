package ticketbusiness

import (
	ticketmodel "cinema/module/ticket/model"
	"context"
)

type UpdateTicketStore interface {
	FindTicket(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*ticketmodel.Ticket, error)

	UpdateTicket(
		ctx context.Context,
		cond map[string]interface{},
		data *ticketmodel.TicketUpdate,
	) error
}

func NewUpdateTicketBiz(store UpdateTicketStore) *updateTicketBiz {
	return &updateTicketBiz{store: store}
}

type updateTicketBiz struct {
	store UpdateTicketStore
}

func (biz *updateTicketBiz) SellTicketToCustomer(ctx context.Context, data *ticketmodel.TicketUpdate) error {
	ticket, err := biz.store.FindTicket(ctx,
		map[string]interface{}{
			"seat_number": data.SeatNumber,
			"show_id":     data.ShowID,
		})

	if err != nil {
		return err
	}

	if ticket.Status == ticketmodel.TicketStatusSold {
		return ticketmodel.ErrTicketHasBeenSold
	}
	if ticket.Status != ticketmodel.TicketStatusAvailable {
		return ticketmodel.ErrTicketUnavailable
	}

	data.Status = ticketmodel.TicketStatusSold

	if err := biz.store.UpdateTicket(ctx, map[string]interface{}{"id": ticket.ID}, data); err != nil {
		return err
	}

	return nil
}
