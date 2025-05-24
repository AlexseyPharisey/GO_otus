package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
		expectError bool
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde", expectError: false},
		{input: "abccd", expected: "abccd", expectError: false},
		{input: "3abc", expected: "", expectError: true},
		{input: "45", expected: "", expectError: true},
		{input: "aaa10b", expected: "", expectError: true},
		{input: "aaa0b", expected: "aab", expectError: false},
		{input: "", expected: "", expectError: false},
		{input: "d\\n5abc", expected: "d\\n\\n\\n\\n\\nabc", expectError: false},
	}

	for _, tt := range tests {
		result, err := UnpackString(tt.input)

		if result != tt.expected && !tt.expectError {
			t.Errorf("INPUT: %q EXPECTED: %q OUTPUT: %q, %q - ERROR", tt.input, tt.expected, result, err)
			continue
		}

		t.Errorf("INPUT: %q EXPECTED: %q OUTPUT: %q, %q - SUCCESS", tt.input, tt.expected, result, err)
	}
}
