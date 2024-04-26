package countrybusiness

import (
	companymodel "cinema/module/company/model"
	"context"
)

type ListCountryStore interface {
	ListCountry(
		context context.Context,
		moreKeys ...string,
	) ([]companymodel.Company, error)
}
type listCountryBusiness struct {
	store ListCountryStore
}

func NewListCountryBusiness(store ListCountryStore) *listCountryBusiness {
	return &listCountryBusiness{store: store}
}
func (business *listCountryBusiness) ListCountry(
	context context.Context,
) ([]companymodel.Company, error) {
	res, err := business.store.ListCountry(context)
	if err != nil {
		return nil, err
	}
	return res, nil
}
