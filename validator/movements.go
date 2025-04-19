package Validator

import (
	"fmt"
	"math"
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

func RotateTwoStacks(a *StackList, b *StackList) {
	RotateStack(a, "")
	RotateStack(b, "")
	Moves = append(Moves, "rr")
}

func SetPrice(a, b *StackList) {
	aLen := a.Length()
	bLen := b.Length()
	current := a.top
	for current != nil {
		current.Cost = current.Index
		if !current.Above {
			current.Cost = aLen - current.Index
		}
		if current.Above {
			current.Cost += current.Target.Index
		} else {
			current.Cost += bLen - current.Target.Index
		}
		current = current.Next
	}
}

func rotate(a *StackList, b *StackList, node *Stack) {
	for b.top.Number != node.Target.Number && a.top.Number != node.Number {
		RotateTwoStacks(a, b)
		a.Index()
		b.Index()
	}
}

// Prep function to position a specific node at the top of a stack
func Prep(sl *StackList, node *Stack, option string) {
	for sl.top.Number != node.Number {
		if option == "a" {
			if node.Above {
				RotateStack(sl, "a")
			} else {
				ReverseRotateStack(sl, "a")
			}
		} else if option == "b" {
			if node.Above {
				RotateStack(sl, "b")
			} else {
				ReverseRotateStack(sl, "b")
			}
		}
	}
}

func Move_a(a *StackList, b *StackList) {
	node := GetCheapest(a)
	if node.Above && node.Target.Above {
		rotate(a, b, node)
	} else if !(node.Above) && !(node.Target.Above) {
		reverse(a, b, node)
	}
	Prep(a, node, "a")
	Prep(b, node.Target, "b")
	PushToStack(b, a, "a")
}

func Move_b(a *StackList, b *StackList) {
	Prep(a, b.top.Target, "a")
	PushToStack(a, b, "b")
}

func SetTargetsA(a *StackList, b *StackList) {
	if a.IsEmpty() || b.IsEmpty() {
		return
	}

	currentA := a.top
	for currentA != nil {
		bestMatchValue := math.MinInt
		var targetNode *Stack

		// Search through the stack for the target
		currentB := b.top
		for currentB != nil {
			// Find closest smaller number
			if currentB.Number < currentA.Number &&
				currentB.Number > bestMatchValue {
				bestMatchValue = currentB.Number
				targetNode = currentB
			}
			currentB = currentB.Next
		}

		// If no smaller number found, set target to max value node
		if bestMatchValue == math.MinInt {
			currentA.Target = FindMaxNode(b)
		} else {
			currentA.Target = targetNode
		}

		currentA = currentA.Next
	}
}

