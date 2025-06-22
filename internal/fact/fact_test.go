package fact

import (
	"testing"
)

type TestCase struct {
	name     string
	input    int
	expected int
}

func TestFact(t *testing.T) {
	testCases := []TestCase{
		{
			name:     "simple",
			input:    5,
			expected: 120,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Fact(tc.input)
			if result != tc.expected {
				t.Errorf("Ожидалось %d, получили %d", tc.expected, result)
			}
		})
	}
}
