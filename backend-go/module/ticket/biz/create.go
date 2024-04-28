package ticketbusiness

import (
	"cinema/common"
	ticketmodel "cinema/module/ticket/model"
	"context"
)

type CreateTicketStore interface {
	Create(context.Context, *ticketmodel.TicketCreate) error
}

type createTicketBusiness struct {
	store CreateTicketStore
}

func NewCreateTicketBusiness(store CreateTicketStore) *createTicketBusiness {
	return &createTicketBusiness{store: store}
}

func (business *createTicketBusiness) CreateTicket(context context.Context, data *ticketmodel.TicketCreate) error {
	if err := business.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(ticketmodel.EntityName, err)
	}
	return nil
}
