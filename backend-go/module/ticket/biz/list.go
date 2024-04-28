package ticketbusiness

import (
	ticketmodel "cinema/module/ticket/model"
	"context"
)

type ListTicketsStore interface {
	ListTicketsWithCondition(
		context context.Context,
		filter *ticketmodel.Filter,
		moreKeys ...string,
	) ([]ticketmodel.Ticket, error)
}
type listTicketsBusiness struct {
	store ListTicketsStore
}

func NewListTicketsBusiness(store ListTicketsStore) *listTicketsBusiness {
	return &listTicketsBusiness{store: store}
}
func (business *listTicketsBusiness) ListTickets(
	context context.Context,
	filter *ticketmodel.Filter,
) ([]ticketmodel.Ticket, error) {
	res, err := business.store.ListTicketsWithCondition(context, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}
