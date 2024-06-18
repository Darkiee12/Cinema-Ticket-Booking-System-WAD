package sdkredis

import (
	"context"
	"flag"
	"github.com/go-redis/redis/v8"
	"log"
)

var (
	defaultUri            = "redis://localhost:6379"
	defaultRedisName      = "DefaultRedis"
	DefaultRedisMaxActive = 0 // 0 is unlimited max active connection
	DefaultRedisMaxIdle   = 10
)

type RedisDBOpt struct {
	Prefix    string
	RedisUri  string
	MaxActive int
	MaxIde    int
}
type redisDB struct {
	name   string
	client *redis.Client
	*RedisDBOpt
}

func NewRedisDB(name string, opts *RedisDBOpt) *redisDB {
	return &redisDB{
		name:       name,
		RedisDBOpt: opts,
	}
}
func (r *redisDB) GetPrefix() string {
	return r.Prefix
}
func (r *redisDB) isDisabled() bool {
	return r.RedisUri == ""
}
func (r *redisDB) InitFlags() {
	prefix := r.Prefix
	if r.Prefix != "" {
		prefix += "-"
	}

	flag.StringVar(&r.RedisUri, prefix+"-uri", defaultUri, "(For go-redis) Redis connection-string. Ex: redis://localhost/0")
	flag.IntVar(&r.MaxActive, prefix+"-pool-max-active", DefaultRedisMaxActive, "(For go-redis) Override redis pool MaxActive")
	flag.IntVar(&r.MaxIde, prefix+"-pool-max-idle", DefaultRedisMaxIdle, "(For go-redis) Override redis pool MaxIdle")
}
func (r *redisDB) Configure() error {
	if r.isDisabled() {
		return nil
	}

	opt, err := redis.ParseURL(r.RedisUri)
	if err != nil {
		return err
	}

	opt.PoolSize = r.MaxActive
	opt.MinIdleConns = r.MaxIde

	client := redis.NewClient(opt)
	// Ping to test Redis connection
	if err := client.Ping(context.Background()).Err(); err != nil {
		return err
	}

	// Connect successfully, assign client to goRedisDB
	r.client = client
	return nil
}
func (r *redisDB) Name() string {
	return r.name
}

func (r *redisDB) Get() interface{} {
	return r.client
}

func (r *redisDB) Run() error {
	return r.Configure()
}
func (r *redisDB) Stop() <-chan bool {
	if r.client != nil {
		if err := r.client.Close(); err != nil {
			log.Println("cannot close ", r.name)
		}
	}

	c := make(chan bool)
	go func() { c <- true }()
	return c
}
