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

// Test function to check for duplicate alues in the stack
func TestHasDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3, 1}, true},
		{[]int{}, false},
		{[]int{0, 0}, true},
		{[]int{-1, -2, -3}, false},
		{[]int{-1, -2, -1}, true},
	}

	for _, tt := range tests {
		result := hasDuplicates(tt.input)
		if result != tt.expected {
			t.Errorf("hasDuplicates(%v) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

// Test function to check if the stack is already sorted
func TestIsSorted(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{}, true},             // empty list is considered sorted
		{[]int{1}, true},            // single element is sorted
		{[]int{1, 2, 3, 4}, true},   // strictly increasing
		{[]int{1, 2, 2, 3}, true},   // non-decreasing
		{[]int{5, 3, 2, 1}, false},  // descending
		{[]int{1, 3, 2}, false},     // out of order in the middle
		{[]int{-5, -2, 0, 3}, true}, // sorted with negatives
		{[]int{-1, -1, -1}, true},   // all equal
	}

	for _, tt := range tests {
		result := isSorted(tt.input)
		if result != tt.expected {
			t.Errorf("isSorted(%v) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

// Test the function to finding the maximum value in the stack
func TestFindMax(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Single element", []int{5}, 5},
		{"All positive", []int{3, 1, 4, 2}, 4},
		{"All negative", []int{-3, -7, -1, -2}, -1},
		{"Mixed values", []int{10, -2, 0, 4}, 10},
		{"Max at end", []int{9, 8, 7, 12}, 12},
		{"Max at start", []int{100, 3, 4, 5}, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMax(tt.input)
			if result != tt.expected {
				t.Errorf("findMax(%v) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}
