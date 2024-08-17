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
		return false, fmt.Errorf("Error")
	} else {
		if !Duplicate(strconv.Itoa(num)) {
			return true, nil
		}
	}
	return false, nil
}

func Duplicate(s string) bool {
	count := 0
	seen := make(map[rune]bool)
	for _, char := range s {
		if seen[char] {
			count++
			return true
		}
		seen[char] = true
	}

	return false
}
