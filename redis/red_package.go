package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const RedPackage = "RedPackage:"
const RedPackageConsume = "RedPackageConsume:"

var redRdb *redis.Client

func main() {
	// 连接redis
	redRdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.10.33:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})

	http.HandleFunc("/sendRedPackage", sendRedPackage)
	http.HandleFunc("/splitRedPackage", splitRedPackage)

	listenErr := http.ListenAndServe(":8080", nil)
	if listenErr != nil {
		panic("http ListenAndServe Fail:" + listenErr.Error())
	}
}

// sendRedPackage 发红包程序
func sendRedPackage(w http.ResponseWriter, r *http.Request) {
	// 用户ID
	uid := r.FormValue("uid")
	// 红包总金额
	totalAmount, _ := strconv.ParseFloat(r.FormValue("total_amount"), 2)
	// 红包个数
	num, _ := strconv.Atoi(r.FormValue("num"))

	if totalAmount <= 0 {
		fmt.Fprint(w, "红包金额必须大于0")
		return
	}
	if num <= 0 {
		fmt.Fprintf(w, "红包个数必须大于0")
		return
	}

	// 将金额转换成分
	totalAmountCent := int(totalAmount * 100)

	// 拆分的红包列表
	var splitRedList []string
	// 已拆红包金额
	useAmount := 0

	// 拆分红包
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		if i == num-1 {
			// 最后一个红包 = 总金额 - 已拆红包金额
			splitRedList = append(splitRedList, strconv.Itoa(totalAmountCent-useAmount))
		} else {
			// 拆分红包金额 = 随机区间（0，（剩余红包金额 / 剩余的人数）* 2）
			max := (totalAmountCent - useAmount) / (num - i) * 2
			splitRedAmount := rand.Intn(max-1) + 1
			if splitRedAmount == 0 {
				fmt.Fprintf(w, "金额太小无法拆分")
				return
			}
			splitRedList = append(splitRedList, strconv.Itoa(splitRedAmount))
			useAmount += splitRedAmount
		}
	}

	ctx := context.Background()
	uuidValue := uuid.New().String()
	// 将拆分的红包列表push到List
	pushErr := redRdb.LPush(ctx, RedPackage+uid+uuidValue, splitRedList).Err()
	if pushErr != nil {
		fmt.Fprintf(w, "Lpush失败：%s\n", pushErr.Error())
		return
	}

	fmt.Fprintf(w, "拆分出来的红包为：%v，单位（分），红包ID：%s\n", splitRedList, uid+uuidValue)
}

// splitRedPackage 抢红包程序
func splitRedPackage(w http.ResponseWriter, r *http.Request) {
	redPackageId := r.FormValue("red_package_id")
	uid := r.FormValue("uid")
	ctx := context.Background()
	// 先查看用户是否已经抢过
	result := redRdb.HGet(ctx, RedPackageConsume+redPackageId, uid).Val()
	if result == "" {
		// 从红包list中pop一个金额
		redAmount, _ := redRdb.LPop(ctx, RedPackage+redPackageId).Int()
		if redAmount == 0 {
			fmt.Fprintf(w, "红包已经被抢光了")
		} else {
			// 抢到红包的记录
			redRdb.HSet(ctx, RedPackageConsume+redPackageId, uid, redAmount)
			fmt.Fprintf(w, "恭喜您抢到了红包：%f\n", float64(redAmount/100))
		}
	} else {
		fmt.Fprintf(w, "您已经抢了该红包")
	}
}
