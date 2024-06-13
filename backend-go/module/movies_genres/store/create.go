package moviesgenresstore

import (
	"cinema/common"
	moviesgenresmodel "cinema/module/movies_genres/model"
	"context"
)

func (store *sqlStore) Create(ctx context.Context, data *moviesgenresmodel.MovieGenre) error {
	if err := store.model(ctx).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
