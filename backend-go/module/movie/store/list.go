package moviestore

import (
	"cinema/common"
	cinemamodel "cinema/module/cinema/model"
	moviemodel "cinema/module/movie/model"
	"context"
)

func (store *sqlStore) ListCompanyWithCondition(
	_ context.Context,
	filter *moviemodel.Filter,
	moreKeys ...string,
) ([]moviemodel.Movie, error) {
	var result []moviemodel.Movie

	db := store.db.Table(cinemamodel.TableName)

	if f := filter; f != nil {
		if f.OwnerID > 0 {
			db = db.Where("user_id = ?", f.OwnerID)
		}
		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Find(&result).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
