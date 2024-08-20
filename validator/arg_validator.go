package validator

import (
	"strconv"
	"strings"
)

func Validate(arg string) bool {
	str := strings.Split(arg, " ")
	seen := make(map[int]struct{})
	for _, v := range str {
		num, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		if _, exists := seen[num]; exists {
			return false
		}
		seen[num] = struct{}{}
	}
	return true
}
