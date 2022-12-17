package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Userinfo struct {
	Uid        int       `orm:"auto"`
	Username   string    `orm:"size(64);null"`
	Department string    `orm:"size(64);null"`
	Created    time.Time `orm:"type(date);null"`
}

type Userdetail struct {
	Uid     int    `orm:"auto"`
	Intro   string `orm:"type(text);null"`
	Profile string `orm:"type(text);null"`
}

func init() {
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(192.168.10.33:3306)/test?charset=utf8", 30)

	// 注册Model
	//orm.RegisterModel(new(Userinfo), new(Userdetail))
	orm.RegisterModel(new(Userinfo))

	// 开启调试
	orm.Debug = true
}

func main() {
	o := orm.NewOrm()

	// 批量插入并返回写入数据
	//uInfos := []Userinfo{
	//	{Username: "shan la", Department: "技术部", Created: time.Date(2022, 12, 17, 0, 0, 0, 0, time.Local)},
	//	{Username: "shan chun", Department: "技术部", Created: time.Date(2022, 12, 18, 0, 0, 0, 0, time.Local)},
	//}
	//for _, uInfo := range uInfos {
	//	id, err := o.Insert(&uInfo)
	//	log.Printf("successNum: %d, ERR: %v\n", id, err)
	//}
	//
	//// 查询后更新数据
	//u1 := Userinfo{Uid: 8}
	//if o.Read(&u1) == nil {
	//	u1.Department = "机动组"
	//	if num, err := o.Update(&u1); err != nil {
	//		fmt.Println("update num: ", num)
	//	}
	//}

	// 根据条件查询数据
	var userinfo Userinfo
	var u2s []*Userinfo

	q := o.QueryTable(userinfo)
	num, err := q.Filter("username", "shan la").All(&u2s)
	fmt.Println("query num:", num, " err: ", err)
	for _, x := range u2s {
		fmt.Println(x)
	}

	// order by
	q1 := o.QueryTable(userinfo)
	num1, err1 := q1.OrderBy("-uid").All(&u2s)
	fmt.Println("query num:", num1, " err: ", err1)
	for _, x := range u2s {
		fmt.Println(x)
	}
}
