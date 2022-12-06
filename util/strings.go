package util

import "strconv"

func MustAtoi(s string) int {
	v, _ := strconv.Atoi(s)

	return v
}
