package util

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
)

func StringToMd5String(str string) string {
	md5str := fmt.Sprintf("%x", StringToMd5Bytes(str))
	return md5str
}

func StringToMd5Bytes(str string) [16]byte {
	data := []byte(str)
	return md5.Sum(data)
}

func IntToMd5Bytes(i int) [16]byte {
	data := IntToBytes(i)
	return md5.Sum(data)
}

//整形转换成字节
func IntToBytes(n int64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
