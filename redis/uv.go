package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123", // no password set
		DB:       0,     // use default DB
	})

	rand.Seed(time.Now().UnixNano())
	fmt.Println("-----模拟有用户点击首页，每个用户来自不同的IP地址，开始-----")
	for i := 0; i < 100; i++ {
		ip := strconv.Itoa(rand.Intn(999)) + "." + strconv.Itoa(rand.Intn(999)) + "." + strconv.Itoa(rand.Intn(999)) + "." + strconv.Itoa(rand.Intn(999))
		hll := rdb.PFAdd(ctx, "hll", ip)
		fmt.Printf("ip=%s, 该IP访问首页的次数：%d\n", ip, hll.Val())
	}
	fmt.Println("-----模拟有用户点击首页，每个用户来自不同的IP地址，结束-----")
}
