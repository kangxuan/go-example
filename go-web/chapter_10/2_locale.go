package main

import (
	"fmt"
	"strconv"
	"time"
)

var locales map[string]map[string]string

func msg(locale, key string) string {
	if v, ok := locales[locale][key]; ok {
		return v
	}
	return ""
}

func date(format string, t time.Time) string {
	year, month, day := t.Date()
	hour, minute, second := t.Clock()

	return fmt.Sprintf(format, year, month, day, hour, minute, second)
}

func moneyFormat(format string, money int64) string {
	return fmt.Sprintf(format, money)
}

func main() {
	// 本地化文本消息
	locales = make(map[string]map[string]string, 2)

	en := make(map[string]string, 10)
	en["pea"] = "pea"
	en["bean"] = "bean"
	en["how old"] = "I am %s years old"
	locales["en"] = en

	zh := make(map[string]string, 10)
	zh["pea"] = "豌豆"
	zh["bean"] = "毛豆"
	zh["how old"] = "我今年%s岁了"
	locales["zh"] = zh

	fmt.Println(msg("zh", "pea"))
	fmt.Printf(msg("en", "how old")+"\n", strconv.Itoa(30))

	// 本地化时间格式
	en["time_zone"] = "America/Chicago"
	zh["time_zone"] = "Asia/Shanghai"

	// 设置时区
	loc, _ := time.LoadLocation(msg("zh", "time_zone"))
	t := time.Now()
	t = t.In(loc)
	fmt.Println(t.Format(time.RFC3339))

	en["date_format"] = "%d-%d-%d %d:%d:%d"
	zh["date_format"] = "%d年%d月%d日 %d时%d分%d秒"

	fmt.Println(date(msg("zh", "date_format"), t))

	// 货币本地化
	en["money"] = "USD %d"
	zh["money"] = "￥%d元"

	fmt.Println(moneyFormat(msg("zh", "money"), 30))

	// 本地化试图和资源，和上面类似
	//
	//views
	//	|--en  // 英文模板
	//		|--images     // 存储图片信息
	//		|--js         // 存储 JS 文件
	//		|--css        // 存储 css 文件
	//		index.tpl     // 用户首页
	//		login.tpl     // 登陆首页
	//	|--zh-CN // 中文模板
	//		|--images
	//		|--js
	//		|--css
	//		index.tpl
	//		login.tpl

	//lang := "zh"
	//s1, _ := template.ParseFiles("views/" + zh + "index.tpl")
	//var VV struct {
	//	Lang string
	//}
	//VV.Lang = lang
	//s1.Execute(os.Stdout, VV)
}
