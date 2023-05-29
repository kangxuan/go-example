package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123", // no password set
		DB:       0,     // use default DB
	})

	// uv
	key := "page:001"
	rdb.Set(ctx, key, 1, -1)

	for i := 0; i < 100; i++ {
		rdb.Incr(ctx, key)
	}

	// 查询uv
	uv := rdb.Get(ctx, key)
	fmt.Printf("uv:%s\n", uv)
}
