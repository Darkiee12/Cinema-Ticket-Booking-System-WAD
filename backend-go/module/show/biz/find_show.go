package showbusiness

import (
	showmodel "cinema/module/show/model"
	"context"
)

type FindShowStore interface {
	FindShow(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*showmodel.Show, error)
}

func NewFindShowBiz(store FindShowStore) *findShowBiz {
	return &findShowBiz{store: store}
}

type findShowBiz struct {
	store FindShowStore
}

func (biz *findShowBiz) FindShowById(ctx context.Context, id int) (*showmodel.Show, error) {
	data, err := biz.store.FindShow(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return data, nil
}
