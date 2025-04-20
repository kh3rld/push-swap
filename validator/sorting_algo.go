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

