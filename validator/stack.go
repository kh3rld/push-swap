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
