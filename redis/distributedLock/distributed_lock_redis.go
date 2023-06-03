package distributedLock

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var rdb *redis.Client
var lockKey string
var expireTime time.Duration

type RedisDistributedLock struct {
}

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.10.33:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
	lockKey = "inventoryRedisLock"
	expireTime = 20 * time.Second
}

// Lock 加锁
func (r *RedisDistributedLock) Lock(ctx context.Context, uuidValue string) error {
	var isLock bool
	luaScript := "if redis.call('exists',KEYS[1]) == 0 or redis.call('hexists',KEYS[1],ARGV[1]) == 1 then " +
		"redis.call('hincrby',KEYS[1],ARGV[1],1) " +
		"redis.call('expire',KEYS[1],ARGV[2]) " +
		"return 1 " +
		"else " +
		"return 0 " +
		"end"
	for {
		isLock, _ = rdb.Eval(ctx, luaScript, []string{lockKey}, uuidValue, expireTime).Bool()
		if isLock == false {
			// 重试等待
			fmt.Println("正在重试加锁：" + uuidValue)
			time.Sleep(20 * time.Millisecond)
		} else {
			break
		}
	}

	// 启动一个定时器，间隔时间查询是否存在，存在则进行延时，如果不存在则取消定时器。结束
	timer := time.NewTicker(expireTime / 3)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-timer.C:
				expireLuaScript := "if redis.call('HEXISTS',KEYS[1],ARGV[1]) == 1 then " +
					"return redis.call('expire',KEYS[1],ARGV[2]) " +
					"else " +
					"return 0 " +
					"end"
				isDelete, _ := rdb.Eval(ctx, expireLuaScript, []string{lockKey}, uuidValue, expireTime).Bool()
				if isDelete == false {
					done <- true
				}
			}
		}
	}()
	return nil
}

// UnLock 解锁
func (r *RedisDistributedLock) UnLock(ctx context.Context, uuidValue string) error {
	luaScript := "if redis.call('HEXISTS',KEYS[1],ARGV[1]) == 0 then " +
		"return nil " +
		"elseif redis.call('HINCRBY',KEYS[1],ARGV[1],-1) == 0 then " +
		"return redis.call('del',KEYS[1]) " +
		"else " +
		"return 0 " +
		"end"
	fmt.Println("解锁：" + uuidValue)
	if rdb.Eval(ctx, luaScript, []string{lockKey}, uuidValue).Val() == nil {
		return errors.New("没有这个锁，HEXISTS查询无")
	}
	return nil
}
