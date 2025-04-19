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

func RotateStack(a *StackList, option string) {
	var arr []int
	len := a.Length() - 1
	temp, _ := a.Pop()
	for i := 0; i <= len-1; i++ {
		value, _ := a.Pop()
		arr = append(arr, value)
	}
	a.Push(temp)
	for i := len - 1; i >= 0; i-- {
		a.Push(arr[i])
	}
	if option == "a" {
		Moves = append(Moves, "ra")
	} else if option == "b" {
		Moves = append(Moves, "rb")
	}
}

func ReverseRotateStack(a *StackList, option string) {
	var arr []int
	len := a.Length() - 1
	for i := 0; i <= len-1; i++ {
		value, _ := a.Pop()
		arr = append(arr, value)
	}
	temp, _ := a.Pop()
	for i := len - 1; i >= 0; i-- {
		a.Push(arr[i])
	}
	a.Push(temp)
	if option == "a" {
		Moves = append(Moves, "rra")
	} else if option == "b" {
		Moves = append(Moves, "reb")
	}
}

func Rrr(a *StackList, b *StackList) {
	ReverseRotateStack(a, "")
	ReverseRotateStack(b, "")
	Moves = append(Moves, "rrr")
}

func reverse(a *StackList, b *StackList, node *Stack) {
	for b.top.Number != node.Target.Number && a.top.Number != node.Number {
		Rrr(a, b)
		a.Index()
		b.Index()
	}
}
