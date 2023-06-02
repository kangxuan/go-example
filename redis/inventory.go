package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strconv"
	"sync"
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
	http.ListenAndServe(":"+port, nil)
}

func saleHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	key := "inventory:001"
	message := sale(ctx, key)
	fmt.Fprintf(w, "端口号："+port+"，"+message)
}

func sale(ctx context.Context, key string) string {
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
}
