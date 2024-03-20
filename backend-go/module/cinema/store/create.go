package cinemastore

import (
	"cinema/common"
	cinemamodel "cinema/module/cinema/model"
	"context"
)

func (s *sqlStore) Create(context context.Context, data *cinemamodel.CinemaCreate) error {
	//data.PrepareForInsert()
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
