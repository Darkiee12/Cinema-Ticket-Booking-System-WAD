package companybusiness

import (
	"cinema/common"
	companymodel "cinema/module/company/model"
	"context"
)

type ListCompanyStore interface {
	ListCompanyWithCondition(
		context context.Context,
		filter *companymodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]companymodel.Company, error)
}
type listCinemaBusiness struct {
	store ListCompanyStore
}

func NewListCompanyBusiness(store ListCompanyStore) *listCinemaBusiness {
	return &listCinemaBusiness{store: store}
}
func (business *listCinemaBusiness) ListCompany(
	context context.Context,
	filter *companymodel.Filter,
	paging *common.Paging,
) ([]companymodel.Company, error) {
	res, err := business.store.ListCompanyWithCondition(context, filter, paging, "Owner")
	if err != nil {
		return nil, err
	}
	return res, nil
}
