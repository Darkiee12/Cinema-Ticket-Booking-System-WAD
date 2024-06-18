package caching

import (
	"context"
	"time"
)

type Cache interface {
	SetTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string, value interface{}) error
	Delete(ctx context.Context, key string) error
	//Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd
}

type requestCounter struct {
	Url   string
	Count int
}

type limitRateEngine struct {
	store Cache
}
