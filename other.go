package util

import (
	"math/rand"
	"time"
)

// RealTimeRand 根据实时时间产生int64随机数
func RealTimeRand() int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63()
}
