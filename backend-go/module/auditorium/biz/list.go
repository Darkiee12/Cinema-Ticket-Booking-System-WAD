package auditoriumbusiness

import (
	"cinema/common"
	auditoriummodel "cinema/module/auditorium/model"
	"context"
)

type ListAuditoriumStore interface {
	ListAuditoriumWithCondition(
		context context.Context,
		filter *auditoriummodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]auditoriummodel.Auditorium, error)
}
type listAuditoriumBusiness struct {
	store ListAuditoriumStore
}

func NewListAuditoriumBusiness(store ListAuditoriumStore) *listAuditoriumBusiness {
	return &listAuditoriumBusiness{store: store}
}

func (business *listAuditoriumBusiness) ListAuditorium(
	context context.Context,
	filter *auditoriummodel.Filter,
	paging *common.Paging,
) ([]auditoriummodel.Auditorium, error) {
	res, err := business.store.ListAuditoriumWithCondition(context, filter, paging, "Cinema")
	if err != nil {
		return nil, err
	}
	return res, nil
}
