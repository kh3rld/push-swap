package Validator

import (
	"fmt"
	"strconv"
	"strings"
)

func Validate(s string) (bool, error) {
	str := strings.TrimSpace(s)
	num, err := strconv.Atoi(str)
	if err != nil {
		return false, fmt.Errorf("error converting string to integer: %w", err)
	}
	return !Duplicate(strconv.Itoa(num)), nil
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
