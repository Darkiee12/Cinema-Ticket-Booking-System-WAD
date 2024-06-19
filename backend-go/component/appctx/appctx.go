package appctx

import (
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetSecretKey() string
	GetRedisClient() *redis.Client
	GetTracer() *trace.Tracer
}

type appCtx struct {
	db        *gorm.DB
	secretKey string
	rdClient  *redis.Client
	tracer    *trace.Tracer
}

func NewAppContext(db *gorm.DB, secretKey string, rdClient *redis.Client, tracer *trace.Tracer) *appCtx {
	return &appCtx{
		db:        db,
		secretKey: secretKey,
		rdClient:  rdClient,
		tracer:    tracer,
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
func (ctx *appCtx) GetTracer() *trace.Tracer {
	return ctx.tracer
}
