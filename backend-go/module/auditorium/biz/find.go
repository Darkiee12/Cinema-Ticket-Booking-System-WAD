package auditoriumbusiness

import (
	auditoriummodel "cinema/module/auditorium/model"
	"context"
)

type FindAuditoriumStore interface {
	FindAuditorium(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*auditoriummodel.Auditorium, error)
}

func NewFindAuditoriumBiz(store FindAuditoriumStore) *findAuditoriumBiz {
	return &findAuditoriumBiz{store: store}
}

type findAuditoriumBiz struct {
	store FindAuditoriumStore
}

func (biz *findAuditoriumBiz) FindAuditoriumById(ctx context.Context, id int) (*auditoriummodel.Auditorium, error) {
	data, err := biz.store.FindAuditorium(ctx, map[string]interface{}{"id": id}, "Cinema", "Owner")

	if err != nil {
		return nil, err
	}

	return data, nil
}
