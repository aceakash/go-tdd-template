package fizzbuzz

import "strconv"

func Count(num int) string {
	if num == 0 {
		return ""
	}
	return strconv.Itoa(num)
}
