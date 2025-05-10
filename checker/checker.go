package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var allowedInstructions = map[string]bool{
	"sa":  true,
	"sb":  true,
	"ss":  true,
	"pa":  true,
	"pb":  true,
	"ra":  true,
	"rb":  true,
	"rr":  true,
	"rra": true,
	"rrb": true,
	"rrr": true,
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	stackA, err := parseInput(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	if hasDuplicates(stackA) {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	stackB := []int{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if !allowedInstructions[line] {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
		applyInstruction(line, &stackA, &stackB)
	}

	if isSorted(stackA) && len(stackB) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

// parseInput converts the input string to a slice of integers
func parseInput(input string) ([]int, error) {
	values := strings.Fields(input)
	stack := make([]int, 0, len(values))

	for _, v := range values {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		stack = append(stack, num)
	}
	return stack, nil
}

// hasDuplicates checks for duplicate integers in the slice
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

// isSorted checks if the stack is sorted in ascending order
func isSorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

// applyInstruction executes the given instruction on the stacks
func applyInstruction(instr string, a *[]int, b *[]int) {
	switch instr {
	case "sa":
		swap(a)
	case "sb":
		swap(b)
	case "ss":
		swap(a)
		swap(b)
	case "pa":
		push(b, a)
	case "pb":
		push(a, b)
	case "ra":
		rotate(a)
	case "rb":
		rotate(b)
	case "rr":
		rotate(a)
		rotate(b)
	case "rra":
		reverseRotate(a)
	case "rrb":
		reverseRotate(b)
	case "rrr":
		reverseRotate(a)
		reverseRotate(b)
	}
}

// swap swaps the top two elements of the stack
func swap(stack *[]int) {
	if len(*stack) >= 2 {
		(*stack)[0], (*stack)[1] = (*stack)[1], (*stack)[0]
	}
}

// push moves the top element from src to dst
func push(src *[]int, dst *[]int) {
	if len(*src) == 0 {
		return
	}
	val := (*src)[0]
	*src = (*src)[1:]
	*dst = append([]int{val}, *dst...)
}

// rotate shifts the stack up (first element becomes last)
func rotate(stack *[]int) {
	if len(*stack) == 0 {
		return
	}
	val := (*stack)[0]
	*stack = append((*stack)[1:], val)
}

// reverseRotate shifts the stack down (last element becomes first)
func reverseRotate(stack *[]int) {
	if len(*stack) == 0 {
		return
	}
	lastIdx := len(*stack) - 1
	val := (*stack)[lastIdx]
	*stack = append([]int{val}, (*stack)[:lastIdx]...)
}
