package config

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var (
	ProjectConfig map[string]string
	Path          string
)

// InitConfig 读取key=value类型的配置文件
func InitConfig() map[string]string {
	config := make(map[string]string)

	f, err := os.Open(Path)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	ProjectConfig = config
	return config
}
