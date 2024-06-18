package caching

import (
	"cinema/module/user/usermodel"
	"context"
	"fmt"
	"log"
)

type FindStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type userCaching struct {
	store     Cache
	findStore FindStore
}

func NewUserCaching(store Cache, findStore FindStore) *userCaching {
	return &userCaching{
		store:     store,
		findStore: findStore,
	}
}

func (uc *userCaching) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	userId := conditions["id"].(int)
	key := fmt.Sprintf("user-%d", userId)
	var userInCache *usermodel.User
	err := uc.store.Get(ctx, key, &userInCache)
	if err != nil {
		log.Println("Error get user from cache", err)
	}
	if userInCache != nil {
		return userInCache, nil
	}

	user, err := uc.findStore.FindUser(ctx, conditions, moreInfo...)

	if err != nil {
		return nil, err
	}

	// Update cache
	// Set cache with key is user id, value is user
	// and ttl is a day
	err = uc.store.SetTTL(ctx, key, user, 24*60*60)
	if err != nil {
		return nil, err
	}
	return user, nil
}
