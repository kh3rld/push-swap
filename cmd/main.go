package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"swap/sortStack"
	"swap/validator"
)

type Stack struct {
	A []int
	B []int
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	stack := sortStack.Stack{
		A: []int{},
		B: []int{},
	}

	if !validator.Validate(os.Args[1]) {
		fmt.Println("Error")
		return
	}
	args := strings.Split(os.Args[1], " ")
	for _, v := range args {
		arg, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error")
			return
		}
		stack.A = append(stack.A, arg)
	}

	fmt.Println("initial stack A:", stack.A)
	instructions := sortStack.PushSwap(&stack)

	fmt.Println("Instructions:")

	for _, instruction := range instructions {
		fmt.Println(instruction)
	}
	fmt.Println("Sorted Stack A:", stack.A)
}
