package config

//存储日志信息
type LogConfig struct {
	Dir string
}

var (
	LoggerConfig *LogConfig
)

func init() {
	println("init config...")
	//初始化log配置
	LoggerConfig = &LogConfig{
		Dir: "./log/",
	}
}
