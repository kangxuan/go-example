package main

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
)

func submit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { // GET输出页面
		t, _ := template.ParseFiles("1_form.html")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}

		// 校验必填字段
		name := r.Form.Get("name")
		if len(name) == 0 {
			log.Println("name字段不能为空")
		}

		// 数字
		age, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil && age > 0 && age < 200 {
			log.Println("age字段必须是数字且大于0小于200")
		}

		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
			log.Println("age字段通过正则校验失败")
		}

		// 汉字
		if h, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("hanzi")); !h {
			log.Println("hanzi字段不是汉字")
		}

		// 电子邮件
		if e, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2,4})$`, r.Form.Get("email")); !e {
			log.Println("email字段不是电子邮件")
		}

		// 手机号
		if m, _ := regexp.MatchString(`^(1[3|4|5|7|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
			log.Println("mobile字段不是手机号")
		}

		// 下拉菜单
		f := r.Form.Get("fruit")
		if !checkFruit(f) {
			log.Println("下拉菜单值未定义")
		}

		// 单选按钮
		g := r.Form.Get("gender")
		if !checkGender(g) {
			log.Println("单选框值未定义")
		}

		// 复选框
		i := r.Form["interest"]
		i = append(i, "string")
		if !checkInterest(i) {
			log.Println("复选框值未定义")
		}

		// 身份证号
		c := r.Form.Get("id_card")
		log.Println(c)
		if !checkIdCard(c) {
			log.Println("身份证号不符合规范")
		}
	}
}

func main() {
	http.HandleFunc("/submit", submit)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// checkFruit 检查下拉菜单
func checkFruit(f string) bool {
	fSlice := []string{"apple", "pear", "banana"}
	for _, item := range fSlice {
		if item == f {
			return true
		}
	}
	return false
}

// checkGender 检查单选框
func checkGender(g string) bool {
	gSlice := []string{"1", "2"}
	for _, item := range gSlice {
		if item == g {
			return true
		}
	}
	return false
}

// checkInterest 检查复选框
func checkInterest(interest []string) bool {
	iSlice := []string{"football", "basketball", "tennis"}
	var tmp bool

	for _, v := range interest {
		tmp = false
		for _, v1 := range iSlice {
			if v == v1 { // 相同则跳过
				tmp = true
				break
			}
		}
		if tmp == false {
			return false
		}
	}

	return true
}

// checkIdCard 检查身份证号
func checkIdCard(idCard string) bool {
	c, _ := regexp.MatchString(`^(\d{15})$`, idCard)
	c1, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, idCard)
	if !c && !c1 {
		return false
	}
	return true
}
