package showbusiness

import (
	showmodel "cinema/module/show/model"
	"context"
)

type GetShowsStore interface {
	GetShows(
		context context.Context,
		id []int,
	) ([]showmodel.Show, error)
}
type getShowsBusiness struct {
	store GetShowsStore
}

func NewGetShowsBusiness(store GetShowsStore) *getShowsBusiness {
	return &getShowsBusiness{store: store}
}
func (business *getShowsBusiness) GetShows(
	context context.Context,
	id []int,
) ([]showmodel.Show, error) {
	res, err := business.store.GetShows(context, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
