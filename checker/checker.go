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

	"swap/sortStack"
	"swap/validator"
)

func executeInstructions(a, b *sortStack.Stack, instructions []string) bool {
	for _, instruction := range instructions {
		switch instruction {
		case "pa":
			sortStack.Pa(a, b)
		case "pb":
			sortStack.Pb(a, b)
		case "sa":
			sortStack.Sa(a)
		case "sb":
			sortStack.Sb(b)
		case "ss":
			sortStack.Ss(a, b)
		case "ra":
			sortStack.Ra(a)
		case "rb":
			sortStack.Rb(b)
		case "rr":
			sortStack.Rr(a, b)
		case "rra":
			sortStack.Rra(a)
		case "rrb":
			sortStack.Rrb(b)
		case "rrr":
			sortStack.Rrr(a, b)
		default:
			fmt.Fprintln(os.Stderr, "Error")
			return false
		}
	}
	return sortStack.IsSorted(a) && len(b.Data) == 0
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	arg := os.Args[1]

	if !validator.Validate(arg) {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	values := strings.Split(arg, " ")
	a := sortStack.NewStack()
	b := sortStack.NewStack()

	for _, v := range values {
		data, err := strconv.Atoi(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
		a.Push(data)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var instructions []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		instructions = append(instructions, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	if executeInstructions(a, b, instructions) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
