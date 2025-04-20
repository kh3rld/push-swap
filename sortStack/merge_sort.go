package sortStack

type Stack struct {
	A []int
	B []int
}

var instructions []string

func PushSwap(s *Stack) []string {
	// Helper functions for operations
	swap := func(stack *[]int, instruction string) {
		if len(*stack) > 1 {
			(*stack)[0], (*stack)[1] = (*stack)[1], (*stack)[0]
			instructions = append(instructions, instruction)
		}
	}

	rotate := func(stack *[]int, instruction string) {
		if len(*stack) > 0 {
			*stack = append((*stack)[1:], (*stack)[0])
			instructions = append(instructions, instruction)
		}
	}

	reverseRotate := func(stack *[]int, instruction string) {
		if len(*stack) > 0 {
			*stack = append([]int{(*stack)[len(*stack)-1]}, (*stack)[:len(*stack)-1]...)
			instructions = append(instructions, instruction)
		}
	}

	push := func(from *[]int, to *[]int, instruction string) {
		if len(*from) > 0 {
			*to = append([]int{(*from)[0]}, *to...)
			*from = (*from)[1:]
			instructions = append(instructions, instruction)
		}
	}

	// // Function to check if a stack is sorted
	// isSorted := func(stack []int) bool {
	// 	for i := 1; i < len(stack); i++ {
	// 		if stack[i-1] > stack[i] {
	// 			return false
	// 		}
	// 	}
	// 	return true
	// }

	// Function to sort a small stack of 3 elements
	sortThree := func(stack *[]int, prefix string) {
		if len(*stack) < 2 {
			return
		}
		if len(*stack) == 2 && (*stack)[0] > (*stack)[1] {
			swap(stack, prefix+"s")
			return
		}
		if (*stack)[0] > (*stack)[1] && (*stack)[0] > (*stack)[2] {
			rotate(stack, prefix+"r")
		} else if (*stack)[1] > (*stack)[2] {
			reverseRotate(stack, prefix+"rr")
		}
		if (*stack)[0] > (*stack)[1] {
			swap(stack, prefix+"s")
		}
	}

	// Main sorting logic
	if len(s.A) <= 3 {
		sortThree(&s.A, "a")
		return instructions
	}

	// Divide and conquer: Push elements to B
	mid := len(s.A) / 2
	for len(s.A) > mid {
		push(&s.A, &s.B, "pb")
	}

	// Sort both halves
	PushSwap(&Stack{A: s.A, B: []int{}})
	PushSwap(&Stack{A: s.B, B: []int{}})

	// Merge sorted halves back into A
	for len(s.B) > 0 {
		push(&s.B, &s.A, "pa")
	}

	return instructions
}
