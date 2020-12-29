package util

import "strconv"

/*
 * 将十进制数字转化为二进制字符串
 */
func ConvertToBin(num int64) string {
	s := ""
	if num == 0 {
		return "0"
	}
	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// 将数字强制性转化为字符串
		s = strconv.FormatInt(lsb, 10) + s
	}
	return s
}
