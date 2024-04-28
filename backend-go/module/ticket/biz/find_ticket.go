package ticketbusiness

import (
	ticketmodel "cinema/module/ticket/model"
	"context"
)

type FindTicketStore interface {
	FindTicket(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*ticketmodel.Ticket, error)
}

func NewFindTicketBiz(store FindTicketStore) *findTicketBiz {
	return &findTicketBiz{store: store}
}

type findTicketBiz struct {
	store FindTicketStore
}

func (biz *findTicketBiz) FindTicketById(ctx context.Context, id int) (*ticketmodel.Ticket, error) {
	data, err := biz.store.FindTicket(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return data, nil
}
