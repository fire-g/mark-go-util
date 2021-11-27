package logger

import (
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	Logger Log
)

type Log struct {
	file    *os.File    //日志文件对象
	dir     string      //日志存储路径
	Info    *log.Logger //提示日志
	Error   *log.Logger //错误日志
	Fatal   *log.Logger //重大错误日志
	Warning *log.Logger //警告日志
	Debug   *log.Logger //错误日志
}

// Init 初始化日志
// 日志存储地址为LoggerConfig.Dir(默认为./data/log/)
// 必须调用此初始化函数才能使用日志，此函数在已在包init函数中调用
func (l *Log) Init() {
	l.ReloadLoggerConfig()
	l.ReopenFile()

	//自定义日志格式
	l.Info = l.newLog("INFO")
	l.Error = l.newLog("ERROR")
	l.Warning = l.newLog("WARNING")
	l.Debug = l.newLog("DEBUG")
	l.Fatal = l.newLog("FATAL")
	// 定时器;定时更新日志文件路径
	go func() {

	}()
}

func (l *Log) newLog(prefix string) *log.Logger {
	format := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	return log.New(io.MultiWriter(l.file, os.Stderr), prefix+" ", format)
}

// SetLogDir 设置日志存储目录
func (l *Log) SetLogDir(dir string) {
	l.dir = dir
}

// ReloadLoggerConfig 重新加载日志配置文件
func (l *Log) ReloadLoggerConfig() {
	err := os.MkdirAll(l.dir, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func (l *Log) ReopenFile() {
	now := time.Now()
	var err error
	now.Format("")
	//日志输出文件
	l.file, err = os.OpenFile(
		l.dir+
			"log_"+strconv.Itoa(now.Year())+"_"+strconv.Itoa(int(now.Month()))+"_"+strconv.Itoa(now.Day())+".log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Fail to open error logger file:", err)
	}
}
