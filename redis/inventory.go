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

/* V4:改进误删锁问题*/
func sale(ctx context.Context, key string) string {
	lockKey := "inventoryRedisLock"
	// 通过uuid+现成ID
	uuidValue := uuid.New().String() + ":" + strconv.FormatInt(goid.Get(), 10)
	// 没有抢占到锁会一直循环获取，这里新增了锁的过期时间，避免程序挂了导致其他正常的程序获取不到锁的问题，这个时间需要根据实际业务估算，且不能使用一条语句保证原子性
	// 查看SexNX源码里层是原子性的。
	for !rdb5.SetNX(ctx, lockKey, uuidValue, 20*time.Second).Val() {
		// 暂停20毫秒继续获取锁
		time.Sleep(20 * time.Millisecond)
	}
	// 使用完成之后删除锁，且要判断误删锁
	defer func() {
		if rdb5.Get(ctx, lockKey).Val() == uuidValue {
			rdb5.Del(ctx, lockKey)
		}
	}()

	inventoryNum, _ := strconv.Atoi(rdb5.Get(ctx, key).Val())
	if inventoryNum > 0 {
		inventoryNum = inventoryNum - 1
		rdb5.Set(ctx, key, inventoryNum, 0)
		return fmt.Sprintf("成功卖出一个商品，剩余库存为：%d\n", inventoryNum)
	} else {
		return fmt.Sprintln("商品已售罄")
	}
}

/* V3:增加分布式锁的过期时间避免程序挂了的问题 */
/*func sale(ctx context.Context, key string) string {
	lockKey := "inventoryRedisLock"
	// 通过uuid+现成ID
	uuidValue := uuid.New().String() + ":" + strconv.FormatInt(goid.Get(), 10)
	// 没有抢占到锁会一直循环获取，这里新增了锁的过期时间，避免程序挂了导致其他正常的程序获取不到锁的问题，这个时间需要根据实际业务估算，且不能使用一条语句保证原子性
	// 查看SexNX源码里层是原子性的。
	for !rdb5.SetNX(ctx, lockKey, uuidValue, 20*time.Second).Val() {
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
}*/

/* V2:分布式锁，问题：设置锁没有过期时间，一旦一台程序挂了且没有释放锁，其他程序就无法抢占锁 */
/*func sale(ctx context.Context, key string) string {
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
}*/

/* v1单机锁版本 问题：分布式系统会超卖 */
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
