package showrepository

import (
	"cinema/common"
	"cinema/component/transactor"
	auditoriumseatsmodel "cinema/module/auditorium_seat/model"
	showmodel "cinema/module/show/model"
	ticketmodel "cinema/module/ticket/model"
	"context"
	"errors"
)

type CreateShowStore interface {
	Create(context context.Context, data *showmodel.ShowCreate) error
	transactor.Transactor
}

type CountShowStore interface {
	Count(ctx context.Context, condition map[string]interface{}) (int64, error)
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
	countShowStore           CountShowStore
}

func NewCreateShowRepo(
	store CreateShowStore,
	ticketStore CreateTicketStore,
	auditoriumSeatsStore ListAuditoriumSeatsStore,
	countShowStore CountShowStore,
) *createShowRepo {
	return &createShowRepo{
		store:                    store,
		createTicketStoreStore:   ticketStore,
		listAuditoriumSeatsStore: auditoriumSeatsStore,
		countShowStore:           countShowStore,
	}
}

func (repo *createShowRepo) CreateShow(
	ctx context.Context,
	data *showmodel.ShowCreate,
) error {
	return repo.store.WithinTransaction(ctx, func(TxCtx context.Context) error {
		condition := map[string]interface{}{
			"date":          data.Date,
			"start_time":    data.StartTime,
			"auditorium_id": data.AuditoriumID,
			"end_time":      data.EndTime,
		}
		count, err := repo.countShowStore.Count(TxCtx, condition)
		if err != nil {
			return common.ErrDB(err)
		}
		if count > 0 {
			return errors.New("there is an another show at this time")
		}
		if data.AuditoriumID == 0 {
			return errors.New("auditorium id is required")
		}
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
