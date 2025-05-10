package Validator

import (
	"fmt"
	"strconv"
	"strings"
)

// parseArgs validates and parses the command-line argument,
// returning a slice of unique integers in reverse order.
func ParseArgs(args []string) ([]int, error) {
	if len(args) > 1 {
		return nil, fmt.Errorf("Error")
	}
	if len(args) == 0 || args[0] == "" {
		return nil, nil
	}

	input := args[0]
	parts := strings.Split(input, " ")
	seen := make(map[int]bool)
	result := make([]int, 0, len(parts))

	for i := len(parts) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(parts[i])
		if err != nil {
			return nil, fmt.Errorf("Error")
		}
		if seen[num] {
			return nil, fmt.Errorf("Error")
		}
		seen[num] = true
		result = append(result, num)
	}

	return result, nil
}
