package auditoriumbusiness

import (
	auditoriummodel "cinema/module/auditorium/model"
	"context"
)

type CreateAuditoriumRepo interface {
	CreateAuditorium(
		ctx context.Context,
		data *auditoriummodel.AuditoriumCreate,
	) error
}

type createAuditoriumRepo struct {
	repo CreateAuditoriumRepo
}

func NewCreateAuditoriumBusiness(repo CreateAuditoriumRepo) *createAuditoriumRepo {
	return &createAuditoriumRepo{repo: repo}
}

func (biz *createAuditoriumRepo) CreateAuditorium(ctx context.Context, data *auditoriummodel.AuditoriumCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.CreateAuditorium(ctx, data); err != nil {
		return err
	}
	return nil
}
