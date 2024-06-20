package main

import (
	"cinema/component/appctx"
	"cinema/plugin/caching/sdkredis"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	redisDB := sdkredis.NewRedisDB(
		"MyRedis",
		&sdkredis.RedisDBOpt{
			RedisUri:  fmt.Sprintf("redis://%s:6379", os.Getenv("localhost")),
			MaxActive: sdkredis.DefaultRedisMaxActive,
			MaxIde:    sdkredis.DefaultRedisMaxIdle,
		})
	// Configure and run the Redis client
	if err := redisDB.Run(); err != nil {
		log.Fatalf("Failed to configure Redis: %v", err)
	}
	defer func() {
		<-redisDB.Stop()
	}()
	client := redisDB.Get().(*redis.Client)
	if client == nil {
		log.Fatalln("Failed to get Redis client")
	}
	appContext := appctx.NewAppContext(nil, "", client, nil) // Assume this initializes your AppContext with Redis client
	cache := sdkredis.NewRedisCache(appContext)

	key := "test_key"
	type valueStruct struct {
		Name string
	}
	value := valueStruct{Name: "test_value"}

	// Set a key
	err := cache.Set(ctx, key, value)
	if err != nil {
		fmt.Println("Set error:", err)
		return
	}

	// Get the key
	var result valueStruct
	err = cache.Get(ctx, key, &result)
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	fmt.Println("Get result:", result)

	// Delete the key
	err = cache.Delete(ctx, key)
	if err != nil {
		fmt.Println("Delete error:", err)
		return
	}

	// Try to get the key again
	err = cache.Get(ctx, key, &result)
	if err == nil {
		fmt.Println("Get after delete should have failed but got:", result)
	} else {
		fmt.Println("Get after delete error as expected:", err)
	}
}
