package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

const signKey = "signKey:"

var rdb1 *redis.Client

func init() {
	rdb1 = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123", // no password set
		DB:       0,     // use default DB
	})
}

func main() {
	ctx := context.Background()
	key := signKey + "2023:" + "1"
	// 第一天签到
	rdb1.SetBit(ctx, key, 1, 1)
	// 第二天签到
	rdb1.SetBit(ctx, key, 2, 1)
	// 第三天签到
	rdb1.SetBit(ctx, key, 3, 0)

	// 获取签到情况
	sign1 := rdb1.GetBit(ctx, key, 1).Val()
	sign2 := rdb1.GetBit(ctx, key, 2).Val()
	sign3 := rdb1.GetBit(ctx, key, 3).Val()
	sign4 := rdb1.GetBit(ctx, key, 4).Val()
	fmt.Printf("sign1:%d\n", sign1)
	fmt.Printf("sign2:%d\n", sign2)
	fmt.Printf("sign3:%d\n", sign3)
	fmt.Printf("sign4:%d\n", sign4)
}
