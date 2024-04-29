package auditoriumrepository

import (
	"cinema/common"
	"cinema/component/transactor"
	auditoriummodel "cinema/module/auditorium/model"
	auditoriumseatsmodel "cinema/module/auditorium_seat/model"
	"context"
)

type CreateAuditoriumStore interface {
	Create(context context.Context, data *auditoriummodel.AuditoriumCreate) error
	transactor.Transactor
}

type CreateAuditoriumSeatsStore interface {
	CreateMultiple(context context.Context, data []*auditoriumseatsmodel.AuditoriumSeat) error
}

type createAuditoriumRepo struct {
	store      CreateAuditoriumStore
	seatsStore CreateAuditoriumSeatsStore
}

func NewCreateAuditoriumRepo(
	store CreateAuditoriumStore,
	seatsStore CreateAuditoriumSeatsStore,
) *createAuditoriumRepo {
	return &createAuditoriumRepo{store: store, seatsStore: seatsStore}
}

func (repo *createAuditoriumRepo) CreateAuditorium(
	ctx context.Context,
	data *auditoriummodel.AuditoriumCreate,
) error {
	return repo.store.WithinTransaction(ctx, func(TxCtx context.Context) error {
		nSeats := data.GetSeats()
		if nSeats <= 0 {
			if err := repo.store.Create(TxCtx, data); err != nil {
				return common.ErrCannotCreateEntity(auditoriummodel.EntityName, err)
			}
			return nil
		}

		if err := repo.store.Create(TxCtx, data); err != nil {
			return common.ErrCannotCreateEntity(auditoriummodel.EntityName, err)
		}

		seats := make([]*auditoriumseatsmodel.AuditoriumSeat, nSeats)
		for i := 0; i < nSeats; i++ {
			seats[i] = &auditoriumseatsmodel.AuditoriumSeat{
				Type:         0,
				AuditoriumID: data.GetID(),
				SeatNumber:   i + 1,
			}

		}
		if err := repo.seatsStore.CreateMultiple(TxCtx, seats); err != nil {
			return common.ErrCannotCreateEntity(auditoriumseatsmodel.EntityName, err)
		}
		return nil
	})
}
