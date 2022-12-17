package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// postFile 模拟浏览器上传文件
func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("upload_file", filename) // 将上传的文件数据写入缓冲区
	if err != nil {
		fmt.Println("写入缓冲区失败")
		return err
	}

	fh, err := os.Open(filename) // 打开文件句柄操作
	if err != nil {
		fmt.Println("打开文件错误")
		return err
	}
	defer func(fh *os.File) {
		err := fh.Close()
		if err != nil {
			log.Fatal("关闭文件错误")
		}
	}(fh)

	_, err = io.Copy(fileWriter, fh) // 复制文件到缓冲区
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	err = bodyWriter.Close()
	if err != nil {
		return err
	}

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("关闭文件错误")
		}
	}(resp.Body)

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(resp.Status)
	fmt.Println(string(respBody))
	return nil
}

func main() {
	targetUrl := "http://localhost:9090/upload" // 上传地址
	filename := "./WechatIMG213.jpeg"           // 文件地址
	err := postFile(filename, targetUrl)
	if err != nil {
		log.Fatal("上传文件，Error: ", err)
	}
}
