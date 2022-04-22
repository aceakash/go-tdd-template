package fizzbuzz

import (
	"strconv"
	"strings"
)

func Count(num int) string {
	if num == 0 {
		return ""
	}

	output := []string{}
	for i := 1; i <= num; i++ {
		if i == 3 {
			output = append(output, "fizz")
			continue
		}
		output = append(output, strconv.Itoa(i))
	}
	return strings.Join(output, " ")
}
