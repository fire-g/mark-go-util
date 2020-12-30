package util

import "bytes"

//字节数组合并
func BytesCombine(b ...[]byte) []byte {
	var buffer bytes.Buffer
	for index := 0; index < len(b); index++ {
		buffer.Write(b[index])
	}
	return buffer.Bytes()
}
