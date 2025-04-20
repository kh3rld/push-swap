package Validator

// validate instructions
func IsValideInstructions(instruction string) bool {
	validInstructions := map[string]struct{}{
		"sa": {}, "sb": {}, "ss": {},
		"pa": {}, "pb": {},
		"ra": {}, "rb": {}, "rr": {},
		"rra": {}, "rrb": {}, "rrr": {},
	}
	_, ok := validInstructions[instruction]
	return ok
}

// execute instructions
func ExecuteInstruction(instruction string, a, b *StackList) {
	switch instruction {
	case "sa":
		SwitchFirstTwo(a, "a")
	case "sb":
		SwitchFirstTwo(b, "b")
	case "ss":
		SwitchFirstTwo(a, "a")
		SwitchFirstTwo(b, "b")
	case "pa":
		PushToStack(a, b, "a")
	case "pb":
		PushToStack(b, a, "b")
	case "ra":
		RotateStack(a, "a")
	case "rb":
		RotateStack(b, "b")
	case "rr":
		RotateStack(a, "a")
		RotateStack(b, "b")
	case "rra":
		ReverseRotateStack(a, "a")
	case "rrb":
		ReverseRotateStack(b, "b")
	case "rrr":
		ReverseRotateStack(a, "a")
		ReverseRotateStack(b, "b")
	}
}
