package showbusiness

import (
	showmodel "cinema/module/show/model"
	"context"
)

type ListShowsStore interface {
	ListShowsWithCondition(
		context context.Context,
		filter *showmodel.Filter,
		moreKeys ...string,
	) ([]showmodel.Show, error)
}
type listShowsBusiness struct {
	store ListShowsStore
}

func NewListShowsBusiness(store ListShowsStore) *listShowsBusiness {
	return &listShowsBusiness{store: store}
}
func (business *listShowsBusiness) ListTickets(
	context context.Context,
	filter *showmodel.Filter,
) ([]showmodel.Show, error) {
	res, err := business.store.ListShowsWithCondition(context, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}
