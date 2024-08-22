package sortStack

type Stack struct {
	A []int
	B []int
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return Merge(left, right)
}

func Merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func PushSwap(s *Stack) []string {
	instructions := []string{}

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

	// Function to push the top element of A to B
	pushAtoB := func() {
		if len(s.A) > 0 {
			instructions = append(instructions, "pb")
			s.B = append(s.B, s.A[0])
			s.A = s.A[1:] // Remove the top element from A
		}
	}

	// Function to push the top element of B to A
	pushBtoA := func() {
		if len(s.B) > 0 {
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
			if s.A[i] > s.A[minIndex] {
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

		// Push the minimum element to stack B
		pushAtoB()
	}

	// Now push back from B to A
	for len(s.B) > 0 {
		minIndex := 0
		for i := 1; i < len(s.B); i++ {
			if s.B[i] > s.B[minIndex] {
				minIndex = i
			}
		}

		// Rotate to bring the minimum element to the top
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

	// Final sorting of stack A using the shortest instructions
	for i := 0; i < len(s.A)-1; i++ {
		for j := 0; j < len(s.A)-i-1; j++ {
			if s.A[j] > s.A[j+1] {
				swap(&s.A, "sa") // Swap the first two elements of A
			}
		}
	}

	// Final sorting of stack A using the shortest instructions
	for i := 0; i < len(s.B)-1; i++ {
		for j := 0; j < len(s.B)-i-1; j++ {
			if s.B[j] > s.B[j+1] {
				swap(&s.B, "sb") // Swap the first two elements of A
			}
		}
	}

	// Return the instructions
	return instructions
}
