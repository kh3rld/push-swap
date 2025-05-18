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

func isSorted(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			return false
		}
	}
	return true
}

func solve(s *Stack) []string {
	switch len(s.A) {
	case 0, 1:
		return []string{}
	case 2:
		return sortTwo(s)
	case 3:
		return sortThree(s)
	case 4, 5:
		return sortFive(s)
	case 6:
		return sortSix(s)
	default:
		return optimizedRadixSort(s)
	}
}

func sortTwo(s *Stack) []string {
	if s.A[0] > s.A[1] {
		return []string{"sa"}
	}
	return []string{}
}
