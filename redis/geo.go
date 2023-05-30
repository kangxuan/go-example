package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

const hotelKey = "hotels"

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123", // no password set
		DB:       0,     // use default DB
	})
}

func main() {
	ctx := context.Background()
	// 添加坐标
	geoAdd(ctx, "希尔顿欢朋酒店", 120.122094, 30.345323)
	geoAdd(ctx, "海外海国际酒店", 120.128721, 30.334236)
	// 获取坐标
	pos, _ := geoPos(ctx, "海外海国际酒店")
	fmt.Println(pos)
	// 获取hash
	hashes, _ := geoHash(ctx, "希尔顿欢朋酒店", "海外海国际酒店")
	fmt.Println(hashes)
	// 获取距离
	dist, _ := geoDist(ctx, "希尔顿欢朋酒店", "海外海国际酒店")
	fmt.Println(dist)
	// 获取就近的酒店
	locations, _ := geoRadius(ctx, 120.120388, 30.341701, 2)
	fmt.Println(locations)
}

// geoAdd 添加坐标
func geoAdd(ctx context.Context, hotelName string, lon float64, lat float64) bool {
	rdb.GeoAdd(ctx, hotelKey, &redis.GeoLocation{Longitude: lon, Latitude: lat, Name: hotelName})
	return true
}

// geoPos 获取坐标
func geoPos(ctx context.Context, hotelName string) ([]*redis.GeoPos, error) {
	pos, err := rdb.GeoPos(ctx, hotelKey, hotelName).Result()
	if err != nil {
		return nil, err
	}
	return pos, nil
}

// geoHash 获取hash
func geoHash(ctx context.Context, hotelNames ...string) ([]string, error) {
	hashes, err := rdb.GeoHash(ctx, hotelKey, hotelNames...).Result()
	if err != nil {
		return nil, err
	}
	return hashes, nil
}

// geoDist 获取距离
func geoDist(ctx context.Context, hotelName1 string, hotelName2 string) (float64, error) {
	dist, err := rdb.GeoDist(ctx, hotelKey, hotelName1, hotelName2, "km").Result()
	if err != nil {
		return 0, err
	}
	return dist, nil
}

// geoRadius 根据距离获取坐标
func geoRadius(ctx context.Context, lon float64, lat float64, radius float64) ([]redis.GeoLocation, error) {
	locations, err := rdb.GeoRadius(ctx, hotelKey, lon, lat, &redis.GeoRadiusQuery{
		Radius: radius,
		Unit:   "km",
	}).Result()
	if err != nil {
		return nil, err
	}
	return locations, nil
}
