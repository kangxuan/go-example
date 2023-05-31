package main

import (
	"context"
	"github.com/redis/go-redis/v9"
)

const cacheKeyA = "cache_key:A"
const cacheKeyB = "cache_key:B"

var rdb4 *redis.Client

func init() {
	rdb4 = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123", // no password set
		DB:       0,     // use default DB
	})
}

// getInfo 获取数据
func getInfo(ctx context.Context, id string) string {
	rB := rdb4.Get(ctx, cacheKeyB+id).Val()
	if rB == "" {
		rA := rdb4.Get(ctx, cacheKeyA+id).Val()
		if rA == "" {
			// 数据库查询，如果到了这一步说明多级缓存有问题，应该报警处理
			return ""
		} else {
			return rA
		}
	} else {
		return rB
	}
}

// addInfo 更新缓存，一般是定时任务进行更新，这里是1小时一次
func addInfo(ctx context.Context, id string, value string) {
	// 先更新A
	rdb4.Del(ctx, cacheKeyA)
	rdb4.Set(ctx, cacheKeyA+id, value, 3600)

	// 再更新B，让B的过期时间晚于A，保证在更新缓存A缓存被删除时，在B中能拿到数据
	rdb4.Del(ctx, cacheKeyB)
	rdb4.Set(ctx, cacheKeyB+id, value, 3610)
}
