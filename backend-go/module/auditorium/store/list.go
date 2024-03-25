package auditoriumstore

import (
	"cinema/common"
	auditoriummodel "cinema/module/auditorium/model"
	"context"
)

func (store *sqlStore) ListAuditoriumWithCondition(
	ctx context.Context,
	filter *auditoriummodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]auditoriummodel.Auditorium, error) {
	var result []auditoriummodel.Auditorium

	db := store.db.Table(auditoriummodel.TableName)

	if f := filter; f != nil {
		if f.CinemaID > 0 {
			db = db.Where("cinema_id = ?", f.CinemaID)
		}
		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)
		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(false)
		paging.NextCursor = last.FakeID.String()
	}

	return result, nil
}
