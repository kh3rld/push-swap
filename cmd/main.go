package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	A []int
	B []int
}

func parseInput(input string) (*Stack, error) {
	values := strings.Fields(input)
	if len(values) == 0 {
		return nil, fmt.Errorf("empty input")
	}

	stack := &Stack{
		A: make([]int, 0, len(values)),
		B: make([]int, 0, len(values)),
	}

	for _, v := range values {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		stack.A = append(stack.A, num)
	}
	return stack, nil
}

func hasDuplicates(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}
