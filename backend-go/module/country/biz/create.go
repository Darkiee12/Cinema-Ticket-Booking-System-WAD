package countrybusiness

import (
	"cinema/common"
	cinemamodel "cinema/module/cinema/model"
	countrymodel "cinema/module/country/model"
	"context"
)

type CreateCountryStore interface {
	Create(context.Context, *countrymodel.Country) error
}

type createCountryBusiness struct {
	store CreateCountryStore
}

func NewCreateCountryBusiness(store CreateCountryStore) *createCountryBusiness {
	return &createCountryBusiness{store: store}
}

func (business *createCountryBusiness) CreateCountry(context context.Context, data *countrymodel.Country) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	if err := business.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(cinemamodel.EntityName, err)
	}
	return nil
}
