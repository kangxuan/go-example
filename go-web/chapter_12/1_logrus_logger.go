package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	// 日志格式设置为 JSON
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 输出 stdout，可以是一个文件
	//logrus.SetOutput(os.Stdout)
	logFile, _ := os.OpenFile("./1_access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logrus.SetOutput(logFile)
	// 只记录严重或以上警告
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	logrus.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	logrus.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// 通过日志重用字段，以下在此程序中不执行，因为Fatal阻断
	contextLogger := logrus.WithFields(logrus.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
