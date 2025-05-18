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

func sortThree(s *Stack) []string {
	var instructions []string
	a, b, c := s.A[0], s.A[1], s.A[2]

	if a < b && b < c {
		return []string{}
	}

	if a == max(a, b, c) {
		instructions = append(instructions, "ra")
		rotate(&s.A)
	} else if b == max(a, b, c) {
		instructions = append(instructions, "rra")
		reverseRotate(&s.A)
	}

	if s.A[0] > s.A[1] {
		instructions = append(instructions, "sa")
	}
	return instructions
}

func sortFive(s *Stack) []string {
	var instructions []string
	for i := 0; i < 2; i++ {
		minIndex := findMinIndex(s.A)
		rotateToTop(s, &instructions, minIndex, "A")
		instructions = append(instructions, "pb")
		push(&s.A, &s.B)
	}

	instructions = append(instructions, sortThree(s)...)

	for len(s.B) > 0 {
		maxIndex := findMaxIndex(s.B)
		rotateToTop(s, &instructions, maxIndex, "B")
		instructions = append(instructions, "pa")
		push(&s.B, &s.A)
	}
	return instructions
}

func sortSix(s *Stack) []string {
	var instructions []string
	for i := 0; i < 3; i++ {
		minIndex := findMinIndex(s.A)
		rotateToTop(s, &instructions, minIndex, "A")
		instructions = append(instructions, "pb")
		push(&s.A, &s.B)
	}

	instructions = append(instructions, sortThree(s)...)

	for len(s.B) > 0 {
		maxIndex := findMaxIndex(s.B)
		rotateToTop(s, &instructions, maxIndex, "B")
		instructions = append(instructions, "pa")
		push(&s.B, &s.A)
	}
	return instructions
}

func optimizedRadixSort(s *Stack) []string {
	var instructions []string
	offset := adjustNegatives(s.A)
	maxBits := calculateMaxBits(s.A)
	currentStack := "A"

	for bit := 0; bit < maxBits; bit++ {
		if currentStack == "A" {
			count := len(s.A)
			for i := 0; i < count; i++ {
				num := s.A[0]
				if ((num+offset)>>bit)&1 == 0 {
					instructions = append(instructions, "pb")
					push(&s.A, &s.B)
				} else {
					instructions = append(instructions, "ra")
					rotate(&s.A)
				}
			}
			currentStack = "B"
		} else {
			count := len(s.B)
			for i := 0; i < count; i++ {
				num := s.B[0]
				if ((num+offset)>>bit)&1 == 0 {
					instructions = append(instructions, "pa")
					push(&s.B, &s.A)
				} else {
					instructions = append(instructions, "rb")
					rotate(&s.B)
				}
			}
			currentStack = "A"
		}
	}

	if currentStack == "B" {
		for len(s.B) > 0 {
			instructions = append(instructions, "pa")
			push(&s.B, &s.A)
		}
	}

	correctNegatives(s, &instructions, offset)
	return instructions
}

func adjustNegatives(nums []int) int {
	minVal := findMin(nums)
	if minVal >= 0 {
		return 0
	}
	offset := -minVal
	for i := range nums {
		nums[i] += offset
	}
	return offset
}

func calculateMaxBits(nums []int) int {
	maxVal := findMax(nums)
	bits := 0
	for maxVal > 0 {
		bits++
		maxVal >>= 1
	}
	return bits + 1
}

func correctNegatives(s *Stack, instructions *[]string, offset int) {
	if offset == 0 {
		return
	}
	splitIndex := 0
	for splitIndex < len(s.A) && s.A[splitIndex] >= offset {
		splitIndex++
	}
	if splitIndex > 0 {
		rotateCount := splitIndex
		if rotateCount <= len(s.A)/2 {
			addReverseRotateA(s, instructions, rotateCount)
		} else {
			addRotateA(s, instructions, len(s.A)-rotateCount)
		}
	}
	for i := range s.A {
		s.A[i] -= offset
	}
}
