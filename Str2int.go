package util4s

import (
	"strconv"
)

// Str2Uint 字符串转无符号整型
func Str2Uint(s string) uint {
	u, err := strconv.Atoi(s)
	if err != nil {
		u = 0
	}
	return uint(u)
}
