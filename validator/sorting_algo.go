package Validator

// TinySort handles sorting of a stack with 3 or fewer elements.
// It applies minimal operations to place the smallest at the top
// using switch, rotate, and reverse rotate operations.
func TinySort(a *StackList) {
	if a.Length() <= 1 {
		// Nothing to sort if the stack has 0 or 1 element
		return
	}

	if a.Length() == 2 {
		// Swap the two if they are out of order
		if a.top.Number > a.top.Next.Number {
			SwitchFirstTwo(a, "a")
		}
		return
	}

	// Get the top three elements of the stack
	first := a.top
	second := first.Next
	third := second.Next

	// Check if the first element is the largest
	if first.Number > second.Number && first.Number > third.Number {
		RotateStack(a, "a")
		if second.Number > third.Number {
			SwitchFirstTwo(a, "a")
		}
	} else if second.Number > first.Number && second.Number > third.Number {
		ReverseRotateStack(a, "a")
		if first.Number > third.Number {
			SwitchFirstTwo(a, "a")
		}
	} else if first.Number > second.Number {
		SwitchFirstTwo(a, "a")
	}

	// Ensure the maximum number is placed correctly
	max, _ := a.FindMax()
	if a.top.Number == max {
		RotateStack(a, "a")
	} else if a.top.Next != nil && a.top.Next.Number == max {
		ReverseRotateStack(a, "a")
	}

	// Final adjustment if first two are still out of order
	if a.top.Number > a.top.Next.Number {
		SwitchFirstTwo(a, "a")
	}
}

// Sort_stack sorts the entire stack 'a' using stack 'b' as auxiliary.
// It first pushes two elements to 'b', sorts the remaining in 'a',
// and then reinserts elements from 'b' into their correct positions in 'a'.
func Sort_stack(a *StackList) {
	b := NewStackList()

	// Push the first two smallest elements to 'b' for better control
	if a.Length() > 3 && !a.IsSorted() {
		PushToStack(b, a, "a")
	}
	if a.Length() > 3 && !a.IsSorted() {
		PushToStack(b, a, "a")
	}

	// Continue moving elements from 'a' to 'b' strategically
	for a.Length() > 3 && !a.IsSorted() {
		Set_a(a, b)
		Move_a(a, b)
	}

	// Sort the remaining 3 elements in 'a'
	TinySort(a)

	// Move all elements from 'b' back to 'a' in sorted order
	for b.top != nil {
		Set_b(a, b)
		Move_b(a, b)
	}

	// Final indexing and bring the minimum to the top
	a.Index()
	MinToTop(a)
}

// Set_a prepares stack 'a' to receive elements from stack 'b'
// by setting targets and computing the cheapest move
func Set_a(a *StackList, b *StackList) {
	a.Index()
	b.Index()
	SetTargetsA(a, b)
	SetPrice(a, b)
	a.SetCheapest()
}

// Set_b prepares stack 'b' for pushing an element back to stack 'a'
// by determining its optimal target location in 'a'
func Set_b(a *StackList, b *StackList) {
	a.Index()
	b.Index()
	SetTargetsB(a, b)
}
