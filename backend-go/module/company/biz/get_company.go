package companybusiness

import (
	companymodel "cinema/module/company/model"
	"context"
)

type FindCompanyStore interface {
	FindCompany(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*companymodel.Company, error)
}

func NewFindCompanyBiz(store FindCompanyStore) *findCompanyBiz {
	return &findCompanyBiz{store: store}
}

type findCompanyBiz struct {
	store FindCompanyStore
}

func (biz *findCompanyBiz) FindCompanyById(ctx context.Context, id int) (*companymodel.Company, error) {
	data, err := biz.store.FindCompany(ctx, map[string]interface{}{"id": id}, "Owner")

	if err != nil {
		return nil, err
	}

	return data, nil
}
