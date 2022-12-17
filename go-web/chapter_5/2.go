package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int    `orm:"auto"`
	Name string `orm:"size(100)"`
}

func init() {
	// 默认数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(192.168.10.33:3306)/test?charset=utf8", 30)

	// 注册定义的model
	orm.RegisterModel(new(User))

	// 打开打印调试
	orm.Debug = false
}

func main() {
	o := orm.NewOrm()
	user := User{Name: "shan la"}

	// 插入
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// 更新
	user.Name = "shanla"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// 读取
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// 删除
	num, err = o.Delete(&u)
	fmt.Printf("NUM:%d, ERR: %v\n", num, err)
}
