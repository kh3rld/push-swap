package validator

import (
	"strconv"
	"strings"
)

<<<<<<< HEAD
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
=======
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
			return nil, fmt.Errorf("error: there are duplicates")
		}
		seen[num] = true
		result = append(result, num)
	}

	return result, nil
>>>>>>> 2bb8de6492297fe6b7d4f4bd302776566bf66bd5
}
