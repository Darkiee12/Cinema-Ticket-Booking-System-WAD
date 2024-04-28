package showrepository

import (
	"cinema/common"
	"cinema/component/transactor"
	auditoriumseatsmodel "cinema/module/auditorium_seat/model"
	showmodel "cinema/module/show/model"
	ticketmodel "cinema/module/ticket/model"
	"context"
)

type CreateShowStore interface {
	Create(context context.Context, data *showmodel.Show) error
	transactor.Transactor
}

type ListAuditoriumSeatsStore interface {
	ListSeatsWithCondition(
		ctx context.Context,
		filter *auditoriumseatsmodel.Filter,
		moreKeys ...string,
	) ([]auditoriumseatsmodel.AuditoriumSeat, error)
}

type CreateTicketStore interface {
	Create(context context.Context, data *ticketmodel.TicketCreate) error
}

type createShowRepo struct {
	store                    CreateShowStore
	createTicketStoreStore   CreateTicketStore
	listAuditoriumSeatsStore ListAuditoriumSeatsStore
}

func NewCreateShowRepo(
	store CreateShowStore,
	seatsStore CreateTicketStore,
	auditoriumSeatsStore ListAuditoriumSeatsStore,
) *createShowRepo {
	return &createShowRepo{
		store:                    store,
		createTicketStoreStore:   seatsStore,
		listAuditoriumSeatsStore: auditoriumSeatsStore,
	}
}

func (repo *createShowRepo) CreateShow(
	ctx context.Context,
	data *showmodel.Show,
) error {
	return repo.store.WithinTransaction(ctx, func(TxCtx context.Context) error {
		seats, err := repo.listAuditoriumSeatsStore.ListSeatsWithCondition(TxCtx, &auditoriumseatsmodel.Filter{
			AuditoriumID: int(data.AuditoriumID),
			Status:       []int{auditoriumseatsmodel.SeatStatusActive},
		})
		if err != nil {
			return common.ErrDB(err)
		}
		if err := repo.store.Create(TxCtx, data); err != nil {
			return common.ErrDB(err)
		}
		for i := range seats {
			if err := repo.createTicketStoreStore.Create(TxCtx, &ticketmodel.TicketCreate{
				ShowID:     int64(data.ID),
				SeatNumber: seats[i].SeatNumber,
				Status:     ticketmodel.TicketStatusAvailable,
			}); err != nil {
				return common.ErrDB(err)
			}
		}
		return nil
	})
}
