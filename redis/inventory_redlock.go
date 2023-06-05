package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
	"net/http"
	"strconv"
)

var port1 string
var err1 error
var rdb6 *goredislib.Client

func init() {
	rdb6 = goredislib.NewClient(&goredislib.Options{
		Addr:     "192.168.10.33:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
	flag.StringVar(&port1, "port", "8080", "启动端口")
	flag.Parse()
}

func main() {
	http.HandleFunc("/sale", saleHandlerByRedLock)
	err1 = http.ListenAndServe(":"+port1, nil)
	if err1 != nil {
		return
	}
}

func saleHandlerByRedLock(w http.ResponseWriter, _ *http.Request) {
	ctx := context.Background()
	key := "inventory:001"
	message := saleByRedLock(ctx, key)
	_, err1 = fmt.Fprintf(w, "端口号："+port1+"，"+message)
	if err1 != nil {
		return
	}
}

/* 多机版 */
/*func saleByRedLock(ctx context.Context, key string) string {
	client6381 := goredislib.NewClient(&goredislib.Options{
		Addr:     "172.100.23.28:6381",
		Password: "123",
		DB:       0,
	})
	pool6381 := goredis.NewPool(client6381)

	client6382 := goredislib.NewClient(&goredislib.Options{
		Addr:     "172.100.23.28:6382",
		Password: "123",
		DB:       0,
	})
	pool6382 := goredis.NewPool(client6382)

	client6383 := goredislib.NewClient(&goredislib.Options{
		Addr:     "172.100.23.28:6383",
		Password: "123",
		DB:       0,
	})
	pool6383 := goredis.NewPool(client6383)

	rs := redsync.New(pool6381, pool6382, pool6383)

	mutex := rs.NewMutex("inventoryRedisLock")
	// 加锁
	mutex.LockContext(ctx)
	// 解锁
	defer mutex.UnlockContext(ctx)
	inventoryNum, _ := strconv.Atoi(rdb6.Get(ctx, key).Val())
	if inventoryNum > 0 {
		inventoryNum = inventoryNum - 1
		rdb6.Set(ctx, key, inventoryNum, 0)
		return fmt.Sprintf("成功卖出一个商品，剩余库存为：%d\n", inventoryNum)
	} else {
		return fmt.Sprintln("商品已售罄")
	}
}*/

/* 单机 */
func saleByRedLock(ctx context.Context, key string) string {
	client1 := goredislib.NewClient(&goredislib.Options{
		Addr:     "192.168.10.33:6379",
		Password: "123456",
		DB:       0,
	})
	pool := goredis.NewPool(client1)
	rs := redsync.New(pool)
	mutex := rs.NewMutex("inventoryRedisLock")
	// 加锁
	mutex.LockContext(ctx)
	// 解锁
	defer mutex.UnlockContext(ctx)
	inventoryNum, _ := strconv.Atoi(rdb6.Get(ctx, key).Val())
	if inventoryNum > 0 {
		inventoryNum = inventoryNum - 1
		rdb6.Set(ctx, key, inventoryNum, 0)
		return fmt.Sprintf("成功卖出一个商品，剩余库存为：%d\n", inventoryNum)
	} else {
		return fmt.Sprintln("商品已售罄")
	}
}
