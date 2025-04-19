package Validator

import (
	"fmt"
	"math"
)

// Stack represents a node in the stack
type Stack struct {
	Number int
	Index  int
	Cost   int
	Above  bool
	Cheap  bool
	Target *Stack
	Next   *Stack
	Prev   *Stack
}

// StackList manages stack operations
type StackList struct {
	top    *Stack
	length int
}

// NewStackList creates and returns a new empty stack
func NewStackList() *StackList {
	return &StackList{}
}

// IsEmpty checks if the stack is empty
func (s *StackList) IsEmpty() bool {
	return s.length == 0
}

// Peek returns the top node without removing it
func (s *StackList) Peek() (*Stack, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	return s.top, true
}

// Push adds a new node to the top of the stack
func (s *StackList) Push(num int) *Stack {
	node := &Stack{Number: num}

	if s.top != nil {
		node.Next = s.top
		s.top.Prev = node
	}
	s.top = node
	s.length++

	return node
}

// Pop removes and returns the number from the stack
func (s *StackList) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	num := s.top.Number
	s.top = s.top.Next
	if s.top != nil {
		s.top.Prev = nil
	}
	s.length--

	return num, true
}

// Length returns the current stack length
func (s *StackList) Length() int {
	return s.length
}

// Print prints all numbers in the stack
func (s *StackList) Print() {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	for curr := s.top; curr != nil; curr = curr.Next {
		fmt.Printf("{ Num-> %d } ", curr.Number)
	}
}

// FindMin returns the minimum number in the stack
func (s *StackList) FindMin() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	min := math.MaxInt
	for curr := s.top; curr != nil; curr = curr.Next {
		if curr.Number < min {
			min = curr.Number
		}
	}
	return min, true
}

// FindMinNode returns the node with the minimum number
func FindMinNode(s *StackList) *Stack {
	if s.IsEmpty() {
		return nil
	}

	minNode := s.top
	for curr := s.top; curr != nil; curr = curr.Next {
		if curr.Number < minNode.Number {
			minNode = curr
		}
	}
	return minNode
}

// FindMax returns the maximum number in the stack
func (s *StackList) FindMax() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	max := math.MinInt
	for curr := s.top; curr != nil; curr = curr.Next {
		if curr.Number > max {
			max = curr.Number
		}
	}
	return max, true
}

// FindMaxNode returns the node with the maximum number
func FindMaxNode(s *StackList) *Stack {
	if s.IsEmpty() {
		return nil
	}

	maxNode := s.top
	for curr := s.top; curr != nil; curr = curr.Next {
		if curr.Number > maxNode.Number {
			maxNode = curr
		}
	}
	return maxNode
}

// Index assigns index and above/below midpoint info
func (s *StackList) Index() {
	mid := s.length / 2
	index := 0
	for curr := s.top; curr != nil; curr = curr.Next {
		curr.Index = index
		curr.Above = index <= mid
		index++
	}
}
