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

	UpdateTickets(
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

	return biz.store.UpdateTickets(ctx, map[string]interface{}{"id": ticket.ID}, data)
}

func (biz *updateTicketBiz) SellManyTicketsToCustomer(ctx context.Context, datas []ticketmodel.TicketUpdate) error {
	var ids []int
	for _, data := range datas {
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
		ids = append(ids, ticket.ID)
	}
	var data *ticketmodel.TicketUpdate
	data.UserID = datas[0].UserID
	data.Status = ticketmodel.TicketStatusSold

	return biz.store.UpdateTickets(ctx, map[string]interface{}{"id": ids}, data)
}
