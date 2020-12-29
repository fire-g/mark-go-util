package logger

import (
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	Info    *log.Logger
	Error   *log.Logger
	Fatal   *log.Logger
	Warning *log.Logger
	Config  *LogConfig
)

//存储日志信息
type LogConfig struct {
	Dir string
}

func init() {
	//初始化log配置
	Config = &LogConfig{
		Dir: "./log/",
	}
	err := os.MkdirAll(Config.Dir, 0777)
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	//日志输出文件
	file, err := os.OpenFile(
		Config.Dir+"log_"+strconv.Itoa(now.Year())+"_"+strconv.Itoa(int(now.Month()))+"_"+strconv.Itoa(now.Day())+".log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Fail to open error logger file:", err)
	}
	//自定义日志格式
	Info = log.New(io.MultiWriter(file, os.Stderr), "INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	Fatal = log.New(io.MultiWriter(file, os.Stderr), "FATAL:", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	Warning = log.New(io.MultiWriter(file, os.Stderr), "WARNING:", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}
