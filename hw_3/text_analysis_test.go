package main

import (
	"strings"
	"testing"
)

func TestTextAnalysis(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input: "cat and dog, one dog,two cats and one man",
			expected: []string{
				"and (2)",
				"one (2)",
				"cat (1)",
				"cats (1)",
				"dog, (1)",
				"dog,two (1)",
				"man (1)",
			},
		},
		{
			input: "Word word WORD word.",
			expected: []string{
				"word (2)",
				"Word (1)",
				"WORD (1)",
			},
		},
		{
			input: "a b c d e f g h i j k l m n o p q r s t a b c d e f g h i j",
			expected: []string{
				"a (2)",
				"b (2)",
				"c (2)",
				"d (2)",
				"e (2)",
				"f (2)",
				"g (2)",
				"h (2)",
				"i (2)",
				"j (2)",
			},
		},
	}

	for _, tt := range tests {
		result := TextAnalysis(tt.input)
		lines := strings.Split(strings.TrimSpace(result), "\n")

		t.Logf("INPUT: %s\n", tt.input)
		t.Logf("RESULT:")
		for _, line := range lines {
			t.Logf("%s", line)
		}

		if len(lines) != len(tt.expected) {
			t.Errorf("expected %d lines, input %d", len(tt.expected), len(lines))
			continue
		}
	}
}
