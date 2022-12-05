package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

// Save 将数据写入到文件中
func (p *Page) Save() (err error) {
	return ioutil.WriteFile(p.Title, p.Body, 0666)
}

// Load 读取文件中的数据
func (p *Page) Load(title string) (err error) {
	p.Title = title
	p.Body, err = ioutil.ReadFile(p.Title)
	return
}

func main() {
	// 展示将数据写入到文件中
	page := Page{
		"Page.md",
		[]byte("# Page\n## Section1\nThis is section1."),
	}

	err := page.Save()
	if err != nil {
		fmt.Println("数据存储失败")
		return
	}

	var newPage Page
	err = newPage.Load("Page.md")
	if err != nil {
		return
	}
	fmt.Println(string(newPage.Body))
}

//# Page
//## Section1
//This is section1.
