package showbusiness

import (
	showmodel "cinema/module/show/model"
	"context"
)

type UpdateShowStore interface {
	FindShow(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*showmodel.Show, error)

	UpdateShow(
		ctx context.Context,
		cond map[string]interface{},
		data *showmodel.ShowUpdate,
	) error
}

func NewUpdateShowBiz(store UpdateShowStore) *updateShowBiz {
	return &updateShowBiz{store: store}
}

type updateShowBiz struct {
	store UpdateShowStore
}

func (biz *updateShowBiz) UpdateShowById(ctx context.Context, id int, data *showmodel.ShowUpdate) error {
	_, err := biz.store.FindShow(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if err := biz.store.UpdateShow(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}

	return nil
}
