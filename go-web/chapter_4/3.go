package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

// upload 上传逻辑
func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		currentTime := time.Now().Unix()
		h := md5.New()
		_, err := io.WriteString(h, strconv.FormatInt(currentTime, 10))
		if err != nil {
			return
		}
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("3_upload.html")
		log.Println(t.Execute(w, token))
	} else {
		err := r.ParseMultipartForm(32 << 20) // 解析MultipartForm数据
		if err != nil {
			log.Fatal("解析MultipartForm数据失败")
		}

		file, handler, err := r.FormFile("upload_file") // 获取上传文件数据
		if err != nil {
			log.Fatal("获取文件失败，Error：", err)
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				log.Fatal("关闭文件失败")
			}
		}(file)

		_, err = fmt.Fprintf(w, "%v", handler.Header) // 向客户端输出文件名
		if err != nil {
			return
		}

		f, err := os.OpenFile("./upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 创建临时文件
		if err != nil {
			log.Fatal("创建并打开临时文件失败", err)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Fatal("关闭临时文件失败")
			}
		}(f)

		_, err = io.Copy(f, file) // 将上传文件的内容复制到创建的临时文件中
		if err != nil {
			return
		}
	}
}

func main() {
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
