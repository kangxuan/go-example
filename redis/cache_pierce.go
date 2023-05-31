package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var rdb3 *redis.Client

func init() {
	rdb3 = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123", // no password set
		DB:       0,     // use default DB
	})
}

func main() {
	ctx := context.Background()
	key := "user:00xxxxxxxx"
	userInfo := getUserInfo(ctx, key)
	fmt.Println(userInfo)
}

func getUserInfo(ctx context.Context, key string) string {
	// 先判断存在redis中
	rValue := rdb3.Get(ctx, key).Val()
	if rValue == "" {
		// redis不存在从查询数据库
		// 这里模拟查询数据库
		daoValue := getDaoValue()
		if daoValue == "" {
			// 采用空对象缓存保证不被恶意的缓存穿透
			rdb3.Set(ctx, key, "", 1*time.Second)
			return ""
		} else {
			return daoValue
		}
	}
	return rValue
}

func getDaoValue() string {
	// 这里只是模拟，最好还采用双检一致性策略
	return ""
}
