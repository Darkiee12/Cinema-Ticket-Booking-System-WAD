package showbusiness

import (
	showmodel "cinema/module/show/model"
	"context"
)

type CreateShowRepo interface {
	CreateShow(
		ctx context.Context,
		data *showmodel.Show,
	) error
}

type createShowRepo struct {
	repo CreateShowRepo
}

func NewCreateShowBusiness(repo CreateShowRepo) *createShowRepo {
	return &createShowRepo{repo: repo}
}

func (biz *createShowRepo) CreateShow(ctx context.Context, data *showmodel.Show) error {
	if err := biz.repo.CreateShow(ctx, data); err != nil {
		return err
	}
	return nil
}
