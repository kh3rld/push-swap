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
