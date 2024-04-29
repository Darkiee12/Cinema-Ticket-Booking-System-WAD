package moviestore

import (
	"cinema/common"
	moviemodel "cinema/module/movie/model"
	"context"
)

func (store *sqlStore) ListMoviesWithCondition(
	_ context.Context,
	_ *moviemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]moviemodel.Movie, error) {
	var result []moviemodel.Movie

	db := store.db.Table(moviemodel.TableName)

	//if f := filter; f != nil {
	//	if f.OwnerID > 0 {
	//		db = db.Where("user_id = ?", f.OwnerID)
	//	}
	//	if len(f.Status) > 0 {
	//		db = db.Where("status in (?)", f.Status)
	//	}
	//}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		db = db.Where("imdb_id > ?", paging.FakeCursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("imdb_id ASC").
		Find(&result).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		paging.NextCursor = last.ImdbID
	}

	return result, nil
}
