package companybusiness

import (
	"cinema/common"
	cinemamodel "cinema/module/cinema/model"
	companymodel "cinema/module/company/model"
	"context"
)

type CreateCompanyStore interface {
	Create(context.Context, *companymodel.CompanyCreate) error
}

type createCompanyBusiness struct {
	store CreateCompanyStore
}

func NewCreateCompanyBusiness(store CreateCompanyStore) *createCompanyBusiness {
	return &createCompanyBusiness{store: store}
}

func (business *createCompanyBusiness) CreateCompany(context context.Context, data *companymodel.CompanyCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	if err := business.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(cinemamodel.EntityName, err)
	}
	return nil
}
