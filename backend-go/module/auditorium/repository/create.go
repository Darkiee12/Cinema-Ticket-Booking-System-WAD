package auditoriumrepository

import (
	"cinema/common"
	auditoriummodel "cinema/module/auditorium/model"
	auditoriumseatsmodel "cinema/module/auditorium_seats/model"
	"context"
)

type Transaction interface {
	OpenTransaction(context context.Context)
	CloseTransaction(context context.Context) error
	RollbackTransaction(context context.Context)
}

type CreateAuditoriumStore interface {
	Create(context context.Context, data *auditoriummodel.AuditoriumCreate) error
	Transaction
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
	nSeats := data.GetSeats()
	if nSeats <= 0 {
		if err := repo.store.Create(ctx, data); err != nil {
			return common.ErrCannotCreateEntity(auditoriummodel.EntityName, err)
		}
		return nil
	}
	repo.store.OpenTransaction(ctx)

	defer func() {
		if r := recover(); r != nil {
			repo.store.RollbackTransaction(ctx)
		}
	}()
	if err := repo.store.Create(ctx, data); err != nil {
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
	if err := repo.seatsStore.CreateMultiple(ctx, seats); err != nil {
		return common.ErrCannotCreateEntity(auditoriumseatsmodel.EntityName, err)
	}
	if err := repo.store.CloseTransaction(ctx); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
