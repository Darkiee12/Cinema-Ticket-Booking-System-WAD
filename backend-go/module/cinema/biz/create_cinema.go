package cinemabuisness

import (
	"cinema/common"
	cinemamodel "cinema/module/cinema/model"
	"context"
)

type CreateCinemaStore interface {
	Create(context.Context, *cinemamodel.CinemaCreate) error
}

type createCinemaBusiness struct {
	store CreateCinemaStore
}

func NewCreateCinemaBusiness(store CreateCinemaStore) *createCinemaBusiness {
	return &createCinemaBusiness{store: store}
}

func (business *createCinemaBusiness) CreateCinema(context context.Context, data *cinemamodel.CinemaCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	if err := business.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(cinemamodel.EntityName, err)
	}
	return nil
}
