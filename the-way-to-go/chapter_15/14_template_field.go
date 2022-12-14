package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name                string
	NonExportedAgeField string
}

func main() {
	// 实例化结构
	p := Person{Name: "Mary", NonExportedAgeField: "31"}

	// 创建一个html模板
	t := template.New("hello")
	// 通过解析模板定义字符串，生成模板的内部表示
	t, _ = t.Parse("hello {{.Name}}, your age is {{.NonExportedAgeField}}")

	// 将模板t输出到命令行输出
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}

	//t1 := template.New("your age is ")
	//t1, _ = t.Parse("your age is {{.nonExportedAgeField}}!")
	//// 将模板t输出到命令行输出
	//if err := t1.Execute(os.Stdout, p); err != nil {
	//	fmt.Println("There was an error:", err.Error())
	//}
}

//hello Mary, your age is 31
