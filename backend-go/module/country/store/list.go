package countrystore

import (
	"cinema/common"
	cinemamodel "cinema/module/cinema/model"
	companymodel "cinema/module/company/model"
	"context"
)

func (store *sqlStore) ListCountry(
	ctx context.Context,
	moreKeys ...string,
) ([]companymodel.Company, error) {
	var result []companymodel.Company

	db := store.db.Table(cinemamodel.TableName)
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
