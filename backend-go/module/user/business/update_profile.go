package userbiz

import (
	"cinema/module/user/usermodel"
	"context"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"log"
)

type UpdateUserStore interface {
	UpdateUser(
		ctx context.Context,
		cond map[string]interface{},
		data *usermodel.UserUpdate,
	) error
}

type UserCacheStore interface {
	Delete(ctx context.Context, key string) error
}

func NewUpdateUserBiz(store UpdateUserStore, cacheStore UserCacheStore) *updateUserBiz {
	return &updateUserBiz{
		store:      store,
		cacheStore: cacheStore,
	}
}

type updateUserBiz struct {
	store      UpdateUserStore
	cacheStore UserCacheStore
}

func (biz *updateUserBiz) UpdateProfileById(ctx context.Context, id int, data *usermodel.UserUpdate) error {
	tracer := ctx.Value("tracer").(trace.Tracer)
	ctx1, span := tracer.Start(ctx, "updateUserBiz.UpdateProfileById")
	defer span.End()
	key := fmt.Sprintf("user-%d", id)
	if err := biz.cacheStore.Delete(ctx, key); err != nil {
		log.Println("Error delete user from cache", err)
	}
	return biz.store.UpdateUser(ctx1, map[string]interface{}{"id": id}, data)
}
