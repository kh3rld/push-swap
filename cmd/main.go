package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Stack represents two stacks: A and B
type Stack struct {
	A []int
	B []int
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	// Parse and validate input
	stack, err := parseInput(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	// Check for duplicates
	if hasDuplicates(stack.A) {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	// If already sorted, return without instructions
	if isSorted(stack.A) {
		return
	}
	args := strings.Split(os.Args[1], " ")
	for _, v := range args {
		arg, err := strconv.Atoi(v)
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

// solve generates sorting instructions for the stack
func solve(s *Stack) []string {
	var instructions []string
	if len(s.A) <= 1 {
		return instructions
	}

	if len(s.A) <= 3 {
		return sortSmall(s)
	}

	if len(s.A) <= 5 {
		return sortMedium(s)
	}

	return sortLargeOptimized(s)
}

// sortSmall sorts a small stack of up to 3 elements
func sortSmall(s *Stack) []string {
	var instructions []string
	if len(s.A) <= 1 {
		return instructions
	}

	if len(s.A) == 2 {
		if s.A[0] > s.A[1] {
			instructions = append(instructions, "sa")
		}
		return instructions
	}

	// Sort 3 elements
	max := findMax(s.A)
	if s.A[0] == max {
		instructions = append(instructions, "ra")
		s.A = append(s.A[1:], s.A[0])
	} else if s.A[1] == max {
		instructions = append(instructions, "rra")
		s.A = append([]int{s.A[len(s.A)-1]}, s.A[:len(s.A)-1]...)
	}

	if s.A[0] > s.A[1] {
		instructions = append(instructions, "sa")
	}
	return instructions
}

// findMax returns the maximum value in the slice
func findMax(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// sortMedium sorts a medium stack of up to 5 elements
func sortMedium(s *Stack) []string {
	var instructions []string

	// Push smallest elements to B until 3 remain in A
	for len(s.A) > 3 {
		min := findMin(s.A)
		for s.A[0] != min {
			index := findIndex(s.A, min)
			if index <= len(s.A)/2 {
				instructions = append(instructions, "ra")
				s.A = append(s.A[1:], s.A[0])
			} else {
				instructions = append(instructions, "rra")
				s.A = append([]int{s.A[len(s.A)-1]}, s.A[:len(s.A)-1]...)
			}
		}
		instructions = append(instructions, "pb")
		s.B = append([]int{s.A[0]}, s.B...)
		s.A = s.A[1:]
	}

	// Sort remaining 3 elements in A
	instructions = append(instructions, sortSmall(s)...)

	// Push back elements from B to A in sorted order
	for len(s.B) > 0 {
		maxVal := findMax(s.B)
		index := findIndex(s.B, maxVal)
		for s.B[0] != maxVal {
			if index <= len(s.B)/2 {
				instructions = append(instructions, "rb")
				s.B = append(s.B[1:], s.B[0])
			} else {
				instructions = append(instructions, "rrb")
				s.B = append([]int{s.B[len(s.B)-1]}, s.B[:len(s.B)-1]...)
			}
			index = findIndex(s.B, maxVal)
		}
		instructions = append(instructions, "pa")
		s.A = append([]int{s.B[0]}, s.A...)
		s.B = s.B[1:]
	}

	return instructions
}

// findMin returns the minimum value in the slice
func findMin(nums []int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

// findIndex returns the index of the value in the slice
func findIndex(nums []int, val int) int {
	for i, num := range nums {
		if num == val {
			return i
		}
	}
	return -1
}

// sortLargeOptimized sorts large stacks using a more efficient algorithm
func sortLargeOptimized(s *Stack) []string {
	var instructions []string
	median := findMedian(s.A)
	remaining := len(s.A)
	for remaining > 3 {
		if s.A[0] <= median {
			instructions = append(instructions, "pb")
			s.B = append([]int{s.A[0]}, s.B...)
			s.A = s.A[1:]
		} else {
			instructions = append(instructions, "ra")
			s.A = append(s.A[1:], s.A[0])
		}
		remaining = len(s.A)
	}

	// Sort remaining 3 elements in A
	instructions = append(instructions, sortSmall(s)...)

	// Push back elements from B to A in sorted order
	for len(s.B) > 0 {
		maxVal := findMax(s.B)
		index := findIndex(s.B, maxVal)
		for s.B[0] != maxVal {
			if index <= len(s.B)/2 {
				instructions = append(instructions, "rb")
				s.B = append(s.B[1:], s.B[0])
			} else {
				instructions = append(instructions, "rrb")
				s.B = append([]int{s.B[len(s.B)-1]}, s.B[:len(s.B)-1]...)
			}
			index = findIndex(s.B, maxVal)
		}
		instructions = append(instructions, "pa")
		s.A = append([]int{s.B[0]}, s.A...)
		s.B = s.B[1:]
	}

	return instructions
}

// findMedian returns the median of the slice
func findMedian(nums []int) int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	return sorted[len(sorted)/2]
}
