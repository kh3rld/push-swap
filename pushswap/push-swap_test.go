package main

import (
	"reflect"
	"testing"
)

// Test function to parse input for accurate and require input.
func TestParseInput(t *testing.T) {
	tests := []struct {
		input       string
		wantStackA  []int
		expectError bool
	}{
		{"1 2 3", []int{1, 2, 3}, false},
		{"10 -5 0", []int{10, -5, 0}, false},
		{"", nil, true},
		{"1 2 a", nil, true},
		{"4 5 5", []int{4, 5, 5}, false},
	}

	for _, tt := range tests {
		stack, err := parseInput(tt.input)
		if tt.expectError {
			if err == nil {
				t.Errorf("expected error for input %q, got nil", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("unexpected error for input %q: %v", tt.input, err)
				continue
			}
			if !reflect.DeepEqual(stack.A, tt.wantStackA) {
				t.Errorf("stack.A = %v, want %v", stack.A, tt.wantStackA)
			}
			if len(stack.B) != 0 {
				t.Errorf("stack.B should be empty, got %v", stack.B)
			}
		}
	}
}
