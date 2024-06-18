package appctx

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetSecretKey() string
	GetRedisClient() *redis.Client
}

type appCtx struct {
	db        *gorm.DB
	secretKey string
	rdClient  *redis.Client
}

func NewAppContext(db *gorm.DB, secretKey string, rdClient *redis.Client) *appCtx {
	return &appCtx{
		db:        db,
		secretKey: secretKey,
		rdClient:  rdClient,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}
func (ctx *appCtx) GetSecretKey() string {
	return ctx.secretKey
}
func (ctx *appCtx) GetRedisClient() *redis.Client {
	return ctx.rdClient
}
