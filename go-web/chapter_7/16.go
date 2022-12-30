package main

import (
	"fmt"
	"html/template"
	"os"
)

// 展示嵌套模板
func main() {
	s1, _ := template.ParseFiles("header.html", "content.html", "footer.html")
	_ = s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	_ = s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	_ = s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	_ = s1.Execute(os.Stdout, nil)

}

//<html>
//<head>
//<title>演示标题</title>
//</head>
//<body>
//
//
//
//<html>
//<head>
//<title>演示标题</title>
//</head>
//<body>
//
//<h1>演示嵌套</h1>
//<ul>
//<li>嵌套使用define定义子模板</li>
//<li>调用使用template</li>
//</ul>
//
//</body>
//</html>
//
//
//
//</body>
//</html>
