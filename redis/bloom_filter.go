package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"hash/fnv"
	"math"
)

type BloomFilter struct {
	Key string
}

var rdb2 *redis.Client

func init() {
	rdb2 = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123", // no password set
		DB:       0,     // use default DB
	})
}

func NewBloomFilter(key string) *BloomFilter {
	return &BloomFilter{
		Key: key,
	}
}

// Add 向布隆过滤器添加值
func (bf *BloomFilter) Add(ctx context.Context, value []byte) {
	// 计算hashcode
	hashValue := bf.hash(value)
	// 通过hashValue和2的32次方取余后，获得对应的下标坑位
	index := hashValue % uint64(math.Pow(2, 32))
	// 设置对应的坑位为1
	rdb2.SetBit(ctx, bf.Key, int64(index), 1)
}

// Contains 判断值是否存在
func (bf *BloomFilter) Contains(ctx context.Context, value []byte) bool {
	// 计算value值的hashcode
	hashValue := bf.hash(value)
	// 通过hashValue和2的32次方取余，获取得到对应的下标坑位
	index := hashValue % uint64(math.Pow(2, 32))
	// 获取对应的坑位值
	if rdb2.GetBit(ctx, bf.Key, int64(index)).Val() == int64(1) {
		return true
	} else {
		return false
	}
}

// hash 通过哈希函数计算出哈希值
func (bf *BloomFilter) hash(value []byte) uint64 {
	hashFunc := fnv.New64()
	hashFunc.Reset()
	hashFunc.Write(value)
	return hashFunc.Sum64()
}

func main() {
	ctx := context.Background()
	filter := NewBloomFilter("whitelistCustomer")

	// 初始化布隆过滤器
	filter.Add(ctx, []byte("customer:001"))
	filter.Add(ctx, []byte("customer:002"))
	filter.Add(ctx, []byte("customer:003"))
	filter.Add(ctx, []byte("customer:004"))

	fmt.Println(filter.Contains(ctx, []byte("customer:001")))
	fmt.Println(filter.Contains(ctx, []byte("customer:002")))
	fmt.Println(filter.Contains(ctx, []byte("customer:003")))
	fmt.Println(filter.Contains(ctx, []byte("customer:004")))
	fmt.Println(filter.Contains(ctx, []byte("customer:005")))

	// 具体的业务
	// 先判断过滤器中是否存在，如果不存在则说明一定不存在，如果存在则有可能存在继续往下走
	if filter.Contains(ctx, []byte("customer:006")) == false {
		fmt.Println("无数据直接返回")
	} else {
		// 过滤器通过查询Redis
		// Redis存在直接返回
		// Redis不存在向数据库查询，此处使用双检一致性算法进一步保证缓存穿透
	}
}
