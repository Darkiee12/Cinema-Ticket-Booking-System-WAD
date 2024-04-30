package moviestore

import (
	"cinema/common"
	moviemodel "cinema/module/movie/model"
	"context"
)

func (store *sqlStore) Create(context context.Context, data *moviemodel.Movie) error {
	//data.PrepareForInsert()
	if err := store.db.Omit("CreatedAt").Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
