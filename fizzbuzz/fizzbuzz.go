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
		switch {
		case i%5 ==0 && i%3 == 0:
			output = append(output, "fizzbuzz")
			continue
		case i%3 == 0:
			output = append(output, "fizz")
			continue
		case i%5 ==0:
			output = append(output, "buzz")
			continue
		}
		output = append(output, strconv.Itoa(i))
	}
	return strings.Join(output, " ")
}
