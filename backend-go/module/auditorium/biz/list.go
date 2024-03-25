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

type ListAuditoriumWithCinemaNameRepo interface {
	ListAuditoriumWithCinemaName(
		context context.Context,
		filter *auditoriummodel.Filter,
		paging *common.Paging,
		cinemaName string,
		moreKeys ...string,
	) ([]auditoriummodel.Auditorium, error)
}

type listAuditoriumWithCinemaNameBusiness struct {
	repo ListAuditoriumWithCinemaNameRepo
}

func NewListAuditoriumWithCinemaNameBusiness(repo ListAuditoriumWithCinemaNameRepo) *listAuditoriumWithCinemaNameBusiness {
	return &listAuditoriumWithCinemaNameBusiness{repo: repo}
}

func (business *listAuditoriumWithCinemaNameBusiness) ListAuditoriumWithCinemaName(
	context context.Context,
	filter *auditoriummodel.Filter,
	paging *common.Paging,
	cinemaName string,
) ([]auditoriummodel.Auditorium, error) {
	res, err := business.repo.ListAuditoriumWithCinemaName(context, filter, paging, cinemaName, "Cinema")
	if err != nil {
		return nil, err
	}
	return res, nil
}
