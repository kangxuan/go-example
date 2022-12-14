package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const lenPath = len("/view/")

var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")
var templates = make(map[string]*template.Template)

type Page struct {
	Title string
	Body  []byte
}

func init() {
	for _, tmpl := range []string{"edit", "view"} {
		templates[tmpl] = template.Must(template.ParseFiles(tmpl + ".html"))
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

// 处理路由
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		title := request.URL.Path[lenPath:]
		if !titleValidator.MatchString(title) {
			http.NotFound(writer, request)
			return
		}
		fn(writer, request, title)
	}
}

// view
func viewHandler(write http.ResponseWriter, request *http.Request, title string) {
	p, err := load(title)
	fmt.Print(title)
	if err != nil {
		http.Redirect(write, request, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(write, "view", p)
}

// edit
func editHandler(write http.ResponseWriter, _ *http.Request, title string) {
	p, err := load(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(write, "edit", p)
}

// save
func saveHandler(writer http.ResponseWriter, request *http.Request, title string) {
	body := request.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(writer, request, "/view/"+title, http.StatusFound)
}

// 写入文件
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// 渲染模板
func renderTemplate(writer http.ResponseWriter, tmpl string, p *Page) {
	err := templates[tmpl].Execute(writer, p)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

// 加载文件内容
func load(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
