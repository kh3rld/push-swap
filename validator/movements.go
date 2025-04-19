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
