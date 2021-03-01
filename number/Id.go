package number

import (
	"math"
	"strings"
)

var (
	chars = "Ab0CdEf2GhIj4KlMn6OpQ8rStUvWxYz" +
		"aB1cDeF3gHiJ5kLmN7oPq9RsTuVwXyZ"
	length = len(chars)
)

//10进制int类型转62进制字符串
func IdToString(num int) string {
	str := ""
	for num > (length - 1) {
		r := num % length
		a := chars[r]
		str = str + string(a)
		num = num / length
	}
	str = str + string(chars[num])
	return str
}

//62进制字符串转10进制int
func StringToId(str string) int {
	s := 0
	for i := 0; i < len(str); i++ {
		j := strings.Index(chars, string(str[i]))
		s = s + j*int(math.Pow(float64(length), float64(i)))
	}
	return s
}
