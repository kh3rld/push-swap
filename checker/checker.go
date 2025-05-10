package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	// Read command line arguments
// 	if len(os.Args) < 2 {
// 		return // No arguments provided, exit without output
// 	}

// 	stackA := make([]int, 0)
// 	seen := make(map[int]bool) // To check for duplicates

// 	// Parse command line arguments
// 	for _, arg := range os.Args[1:] {
// 		num, err := strconv.Atoi(arg)
// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, "Error") // Invalid integer
// 			return
// 		}
// 		if seen[num] {
// 			fmt.Fprintln(os.Stderr, "Error") // Duplicate found
// 			return
// 		}
// 		seen[num] = true
// 		stackA = append(stackA, num)
// 	}

// 	// Read instructions from standard input
// 	scanner := bufio.NewScanner(os.Stdin)
// 	instructions := make([]string, 0)

// 	for scanner.Scan() {
// 		line := strings.TrimSpace(scanner.Text())
// 		if line == "" {
// 			continue // Skip empty lines
// 		}
// 		instructions = append(instructions, line)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Fprintln(os.Stderr, "Error") // Error reading input
// 		return
// 	}

// 	// Execute instructions
// 	stackB := make([]int, 0)

// 	for _, instruction := range instructions {
// 		switch instruction {
// 		case "sa":
// 			if len(stackA) < 2 {
// 				fmt.Fprintln(os.Stderr, "Error") // Not enough elements to swap
// 				return
// 			}
// 			stackA[0], stackA[1] = stackA[1], stackA[0]
// 		case "sb":
// 			if len(stackB) < 2 {
// 				fmt.Fprintln(os.Stderr, "Error") // Not enough elements to swap
// 				return
// 			}
// 			stackB[0], stackB[1] = stackB[1], stackB[0]
// 		case "pa":
// 			if len(stackB) == 0 {
// 				continue // No elements to push
// 			}
// 			stackA = append([]int{stackB[0]}, stackA...) // Push from B to A
// 			stackB = stackB[1:]                          // Remove the pushed element
// 		case "pb":
// 			if len(stackA) == 0 {
// 				continue // No elements to push
// 			}
// 			stackB = append([]int{stackA[0]}, stackB...) // Push from A to B
// 			stackA = stackA[1:]                          // Remove the pushed element
// 		case "ra":
// 			if len(stackA) > 0 {
// 				stackA = append(stackA[1:], stackA[0]) // Rotate A
// 			}
// 		case "rb":
// 			if len(stackB) > 0 {
// 				stackB = append(stackB[1:], stackB[0]) // Rotate B
// 			}
// 		case "rra":
// 			if len(stackA) > 0 {
// 				stackA = append([]int{stackA[len(stackA)-1]}, stackA[:len(stackA)-1]...) // Reverse Rotate A
// 			}
// 		case "rrb":
// 			if len(stackB) > 0 {
// 				stackB = append([]int{stackB[len(stackB)-1]}, stackB[:len(stackB)-1]...) // Reverse Rotate B
// 			}
// 		default:
// 			fmt.Fprintln(os.Stderr, "Error") // Invalid instruction
// 			return
// 		}
// 	}

// 	// Check if stack A is sorted and stack B is empty
// 	if isSorted(stackA) && len(stackB) == 0 {
// 		fmt.Println("OK")
// 	} else {
// 		fmt.Println("KO")
// 	}
// }

// // Helper function to check if a stack is sorted
// func isSorted(stack []int) bool {
// 	for i := 1; i < len(stack); i++ {
// 		if stack[i] < stack[i-1] {
// 			return false
// 		}
// 	}
// 	return true
// }

// package main

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
