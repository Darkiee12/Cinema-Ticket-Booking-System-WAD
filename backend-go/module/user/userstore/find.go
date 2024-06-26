package userstore

import (
	"cinema/common"
	"cinema/module/user/usermodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(_ context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	//_, span := trace.StartSpan(ctx, "List Restaurant Business")
	//defer span.End()

	db := s.db.Table(usermodel.TableName)

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}
