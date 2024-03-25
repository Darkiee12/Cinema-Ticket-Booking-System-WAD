package auditoriumrepository

import (
	"cinema/common"
	auditoriummodel "cinema/module/auditorium/model"
	cinemamodel "cinema/module/cinema/model"
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

type CinemaFindStore interface {
	FindCinema(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*cinemamodel.Cinema, error)
}

type listAuditoriumWithCinemaNameRepo struct {
	store           ListAuditoriumStore
	findCinemaStore CinemaFindStore
}

func NewListAuditoriumWithCinemaNameRepo(
	store ListAuditoriumStore,
	findCinemaStore CinemaFindStore,
) *listAuditoriumWithCinemaNameRepo {
	return &listAuditoriumWithCinemaNameRepo{store: store, findCinemaStore: findCinemaStore}
}

func (repo *listAuditoriumWithCinemaNameRepo) ListAuditoriumWithCinemaName(
	ctx context.Context,
	filter *auditoriummodel.Filter,
	paging *common.Paging,
	cinemaName string,
	moreKeys ...string,
) ([]auditoriummodel.Auditorium, error) {
	cinema, err := repo.findCinemaStore.FindCinema(ctx, map[string]interface{}{"name": cinemaName})
	if err != nil {
		return nil, err
	}

	filter.CinemaID = cinema.ID
	res, err := repo.store.ListAuditoriumWithCondition(ctx, filter, paging, moreKeys...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
