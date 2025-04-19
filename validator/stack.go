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
