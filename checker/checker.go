package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	A []int
	B []int
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	// Parse input stack A
	arg := os.Args[1]
	values := strings.Split(arg, " ")
	stack := &Stack{A: []int{}, B: []int{}}
	for _, v := range values {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
		stack.A = append(stack.A, num)
	}

	// Read instructions from standard input
	scanner := bufio.NewScanner(os.Stdin)
	instructions := []string{}
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	// Execute instructions
	for _, instr := range instructions {
		switch instr {
		case "sa":
			swap(&stack.A)
		case "sb":
			swap(&stack.B)
		case "ss":
			swap(&stack.A)
			swap(&stack.B)
		case "pa":
			push(&stack.B, &stack.A)
		case "pb":
			push(&stack.A, &stack.B)
		case "ra":
			rotate(&stack.A)
		case "rb":
			rotate(&stack.B)
		case "rr":
			rotate(&stack.A)
			rotate(&stack.B)
		case "rra":
			reverseRotate(&stack.A)
		case "rrb":
			reverseRotate(&stack.B)
		case "rrr":
			reverseRotate(&stack.A)
			reverseRotate(&stack.B)
		default:
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
	}

	// Check if stack A is sorted and stack B is empty
	if isSorted(stack.A) && len(stack.B) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

// Helper functions
func swap(stack *[]int) {
	if len(*stack) > 1 {
		(*stack)[0], (*stack)[1] = (*stack)[1], (*stack)[0]
	}
}

func push(from *[]int, to *[]int) {
	if len(*from) > 0 {
		*to = append([]int{(*from)[0]}, *to...)
		*from = (*from)[1:]
	}
}

func rotate(stack *[]int) {
	if len(*stack) > 0 {
		*stack = append((*stack)[1:], (*stack)[0])
	}
}

func reverseRotate(stack *[]int) {
	if len(*stack) > 0 {
		*stack = append([]int{(*stack)[len(*stack)-1]}, (*stack)[:len(*stack)-1]...)
	}
}

func isSorted(stack []int) bool {
	for i := 1; i < len(stack); i++ {
		if stack[i-1] > stack[i] {
			return false
		}
	}
	return true
}
