package cinemabuisness

import (
	cinemamodel "cinema/module/cinema/model"
	"context"
)

type FindCinemaStore interface {
	FindCinema(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*cinemamodel.Cinema, error)
}

func NewFindCinemaBiz(store FindCinemaStore) *findCinemaBiz {
	return &findCinemaBiz{store: store}
}

type findCinemaBiz struct {
	store FindCinemaStore
}

func (biz *findCinemaBiz) FindCinemaById(ctx context.Context, id int) (*cinemamodel.Cinema, error) {
	data, err := biz.store.FindCinema(ctx, map[string]interface{}{"id": id}, "Owner")

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (biz *findCinemaBiz) FindCinemaByName(ctx context.Context, name string) (*cinemamodel.Cinema, error) {
	data, err := biz.store.FindCinema(ctx, map[string]interface{}{"name": name}, "Owner")

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (biz *findCinemaBiz) FindCinemaByOwnerID(ctx context.Context, ownerID int) (*cinemamodel.Cinema, error) {
	data, err := biz.store.FindCinema(ctx, map[string]interface{}{"owner_id": ownerID}, "Owner")

	if err != nil {
		return nil, err
	}

	return data, nil
}
