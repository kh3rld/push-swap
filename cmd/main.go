package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"swap/sortStack"
	"swap/validator"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	arg := os.Args[1]

	if !validator.Validate(arg) {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	// Parse input into stack A
	values := strings.Split(arg, " ")
	stack := &sortStack.Stack{A: []int{}, B: []int{}}
	for _, v := range values {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
		stack.A = append(stack.A, num)
	}

	// Generate instructions using PushSwap
	instructions := sortStack.PushSwap(stack)

	// Output instructions
	for _, instr := range instructions {
		fmt.Println(instr)
	}
}
