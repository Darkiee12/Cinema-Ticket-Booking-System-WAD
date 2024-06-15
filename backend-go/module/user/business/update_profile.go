package userbiz

import (
	"cinema/module/user/usermodel"
	"context"
)

type UpdateUserStore interface {
	UpdateUser(
		ctx context.Context,
		cond map[string]interface{},
		data *usermodel.UserUpdate,
	) error
}

func NewUpdateUserBiz(store UpdateUserStore) *updateUserBiz {
	return &updateUserBiz{store: store}
}

type updateUserBiz struct {
	store UpdateUserStore
}

func (biz *updateUserBiz) UpdateProfileById(ctx context.Context, id int, data *usermodel.UserUpdate) error {
	return biz.store.UpdateUser(ctx, map[string]interface{}{"id": id}, data)
}
