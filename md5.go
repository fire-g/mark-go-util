package util

import (
	"crypto/md5"
	"fmt"
)

func Md5String(str string) string {
	md5str := fmt.Sprintf("%x", Md5Bytes(str))
	return md5str
}

func Md5Bytes(str string) [16]byte {
	data := []byte(str)
	return md5.Sum(data)
}
