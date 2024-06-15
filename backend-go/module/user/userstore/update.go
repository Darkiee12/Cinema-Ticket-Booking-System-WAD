package userstore

import (
	"cinema/common"
	"cinema/module/user/usermodel"
	"context"
)

func (store *sqlStore) UpdateUser(
	_ context.Context,
	cond map[string]interface{},
	data *usermodel.UserUpdate,
) error {
	if err := store.db.Where(cond).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
