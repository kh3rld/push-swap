package Validator

import (
	"strconv"
	"strings"
)

func Validate(s string) bool {
	str := strings.TrimSpace(s)
	num, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return !Duplicate(strconv.Itoa(num))
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
