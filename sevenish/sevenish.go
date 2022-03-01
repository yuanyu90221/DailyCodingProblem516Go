package sevenish

import "strconv"

func FindSevenishNumber(n int64) int64 {
	base2 := strconv.FormatInt(n, 2)
	result, _ := strconv.ParseInt(base2, 7, 64)
	return result
}
