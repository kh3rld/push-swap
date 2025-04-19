package Validator

import (
	"fmt"
)

var Moves []string

func PushToStack(dst, src *StackList, option string) {
	value, ok := src.Pop()
	if !ok {
		fmt.Println("the stack is empty")
		return
	}
	dst.Push(value)
	if option == "a" {
		Moves = append(Moves, "pb")
	} else if option == "b" {
		Moves = append(Moves, "pa")
	}
}

func SwitchFirstTwo(stack *StackList, option string) {
	first, ok := stack.Pop()
	if !ok {
		return
	}
	second, ok := stack.Pop()
	if !ok {
		stack.Push(first)
		return
	}
	stack.Push(first)
	stack.Push(second)
	if option == "a" {
		Moves = append(Moves, "sa")
	} else if option == "b" {
		Moves = append(Moves, "sb")
	}
}

func SwitchBothStacks(a, b *StackList) {
	SwitchFirstTwo(a, "")
	SwitchFirstTwo(b, "")
	Moves = append(Moves, "ss")
}
