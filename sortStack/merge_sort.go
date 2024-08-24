package sortStack

type Stack struct {
	A []int
	B []int
}

var instructions []string

func PushSwap(s *Stack) []string {
	// Helper functions for operations
	swap := func(stack *[]int, instruction string) {
		if len(*stack) >= 2 {
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

	pushAtoB := func() {
		if len(s.A) > 0 && (len(s.B) == 0 || s.A[0] > s.B[len(s.B)-1]) {
			instructions = append(instructions, "pb")
			s.B = append(s.B, s.A[0])
			s.A = s.A[1:]
		}
	}

	// Function to push the top element of B to A
	pushBtoA := func() {
		if len(s.B) > 0 && (len(s.A) == 0 || s.B[0] > s.A[len(s.A)-1]) {
			instructions = append(instructions, "pa")
			s.A = append(s.A, s.B[len(s.B)-1])
			s.B = s.B[:len(s.B)-1] // Remove the top element from B
		}
	}

	// Sort stack A using the defined instructions
	for len(s.A) > 0 {
		// Find the minimum element in stack A
		minIndex := 0
		for i := 1; i < len(s.A); i++ {
			if s.A[i] < s.A[minIndex] {
				minIndex = i
			}
		}

		// Rotate to bring the minimum element to the top
		if minIndex > 0 {
			if minIndex <= len(s.A)/2 {
				for i := 0; i < minIndex; i++ {
					rotate(&s.A, "ra")
				}
			} else {
				for i := 0; i < len(s.A)-minIndex; i++ {
					reverseRotate(&s.A, "rra")
				}
			}
		}
		pushAtoB()
	}

	for len(s.B) > 0 {
		minIndex := 0
		for i := 1; i < len(s.B); i++ {
			if s.B[i] < s.B[minIndex] {
				minIndex = i
			}
		}

		if minIndex > 0 {
			if minIndex <= len(s.B)/2 {
				for i := 0; i < minIndex; i++ {
					rotate(&s.B, "rb")
				}
			} else {
				for i := 0; i < len(s.B)-minIndex; i++ {
					reverseRotate(&s.B, "rrb")
				}
			}
		}

		pushBtoA()
	}
	for len(s.A) > 0 && len(s.B) > 0 {
		reverseRotate(&s.A, "rrr")
		reverseRotate(&s.B, "rrr")
		pushBtoA()
	}
	for i := 0; i < len(s.B)-1; i++ {
		for j := 0; j < len(s.B)-i-1; j++ {
			if s.B[j] > s.B[j+1] {
				swap(&s.B, "sb")
			}
		}
	}
	return instructions
}
