package logger

import (
	"github.com/fire-g/mark-go-util/config"
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
	file    *os.File
	err     error
)

//重新加载日志配置文件
func ReloadLoggerConfig() {
	err = os.MkdirAll(config.LoggerConfig.Dir, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func ReopenFile() {
	now := time.Now()
	var f *os.File
	//日志输出文件
	f, err = os.OpenFile(
		config.LoggerConfig.Dir+
			"log_"+strconv.Itoa(now.Year())+"_"+strconv.Itoa(int(now.Month()))+"_"+strconv.Itoa(now.Day())+".log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Fail to open error logger file:", err)
	}
	file = f
}

func init() {
	ReloadLoggerConfig()
	ReopenFile()

	//自定义日志格式
	Info = log.New(io.MultiWriter(file, os.Stderr), "INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	Fatal = log.New(io.MultiWriter(file, os.Stderr), "FATAL:", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	Warning = log.New(io.MultiWriter(file, os.Stderr), "WARNING:", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}
