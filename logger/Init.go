package logger

func init() {
	//初始化日志对象
	Logger.SetLogDir("./data/log/")
	Logger.Init()
}
