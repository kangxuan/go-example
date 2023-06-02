package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/petermattis/goid"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var rdb5 *redis.Client
var lock sync.Mutex

var port string

func init() {
	rdb5 = redis.NewClient(&redis.Options{
		Addr:     "192.168.10.33:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
	flag.StringVar(&port, "port", "8080", "启动端口")
	flag.Parse()
}

func main() {
	http.HandleFunc("/sale", saleHandler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return
	}
}

func saleHandler(w http.ResponseWriter, _ *http.Request) {
	ctx := context.Background()
	key := "inventory:001"
	message := sale(ctx, key)
	_, err := fmt.Fprintf(w, "端口号："+port+"，"+message)
	if err != nil {
		return
	}
}

func sale(ctx context.Context, key string) string {
	lockKey := "inventoryRedisLock"
	// 通过uuid+现成ID
	uuidValue := uuid.New().String() + ":" + strconv.FormatInt(goid.Get(), 10)
	// 没有抢占到锁会一直循环获取
	for !rdb5.SetNX(ctx, lockKey, uuidValue, 0).Val() {
		// 暂停20毫秒继续获取锁
		time.Sleep(20 * time.Millisecond)
	}
	// 使用完成之后删除锁
	defer rdb5.Del(ctx, lockKey)

	inventoryNum, _ := strconv.Atoi(rdb5.Get(ctx, key).Val())
	if inventoryNum > 0 {
		inventoryNum = inventoryNum - 1
		rdb5.Set(ctx, key, inventoryNum, 0)
		return fmt.Sprintf("成功卖出一个商品，剩余库存为：%d\n", inventoryNum)
	} else {
		return fmt.Sprintln("商品已售罄")
	}
}

/* v1单机锁版本 */
/*func sale(ctx context.Context, key string) string {
	lock.Lock()
	defer lock.Unlock()
	inventoryNum, _ := strconv.Atoi(rdb5.Get(ctx, key).Val())
	if inventoryNum > 0 {
		inventoryNum = inventoryNum - 1
		rdb5.Set(ctx, key, inventoryNum, 0)
		return fmt.Sprintf("成功卖出一个商品，剩余库存为：%d\n", inventoryNum)
	} else {
		return fmt.Sprintln("商品已售罄")
	}
}*/
