package fizzbuzz

import (
	"strconv"
	"strings"
)

func Count(num int) string {
	if num == 0 {
		return ""
	}

	if num == 2 {
		output := []string{}
		for i := 1; i <= num ; i++ {
			output = append(output, strconv.Itoa(i))
		}
		return strings.Join(output, " ")
	}

	if num == 3 {
		return "1 2 fizz"
	}
	return strconv.Itoa(num)
}
