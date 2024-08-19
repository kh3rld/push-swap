package validator

import (
	"strconv"
	"strings"
)

func Validate(s string) bool {
	str := strings.Fields(s)
	var b bool
	for _, x := range str {
		num, err := strconv.Atoi(x)
		if err != nil {
			return false
		}
		b = !Duplicate(strconv.Itoa(num))
	}
	return b
}

func Duplicate(s string) bool {
	seen := make(map[rune]bool)
	for _, c := range s {
		if seen[c] {
			return true
		}
		seen[c] = true
	}
	return false
}
