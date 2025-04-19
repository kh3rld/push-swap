package Validator

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
