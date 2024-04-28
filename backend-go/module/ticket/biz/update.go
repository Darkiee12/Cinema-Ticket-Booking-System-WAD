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

func (biz *updateTicketBiz) SellTicketToCustomer(ctx context.Context, seatNumber int, showID, userID int64) error {
	ticket, err := biz.store.FindTicket(ctx, map[string]interface{}{"seat_number": seatNumber, "show_id": showID})

	if err != nil {
		return err
	}

	if ticket.Status == ticketmodel.TicketStatusSold {
		return ticketmodel.ErrTicketHasBeenSold
	}
	if ticket.Status != ticketmodel.TicketStatusAvailable {
		return ticketmodel.ErrTicketUnavailable
	}

	data := &ticketmodel.TicketUpdate{
		Status: ticketmodel.TicketStatusSold,
		UserID: userID,
	}

	if err := biz.store.UpdateTicket(ctx, map[string]interface{}{"id": ticket.ID}, data); err != nil {
		return err
	}

	return nil
}
