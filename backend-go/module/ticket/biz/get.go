package ticketbusiness

import (
	ticketmodel "cinema/module/ticket/model"
	"context"
)

type GetTicketsStore interface {
	GetTickets(
		context context.Context,
		id []int,
	) ([]ticketmodel.Ticket, error)
}
type getTicketsBusiness struct {
	store GetTicketsStore
}

func NewGetTicketsBusiness(store GetTicketsStore) *getTicketsBusiness {
	return &getTicketsBusiness{store: store}
}
func (business *getTicketsBusiness) GetTickets(
	context context.Context,
	id []int,
) ([]ticketmodel.Ticket, error) {
	res, err := business.store.GetTickets(context, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
