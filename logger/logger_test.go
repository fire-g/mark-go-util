package logger

import "testing"

func TestLog_Init(t *testing.T) {
	Logger.Info.Println("测试日志")
}
