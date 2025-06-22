package sumarray

import "testing"

type TestCase struct {
	name     string
	input    []int
	expected int
}

func TestSumArray(t *testing.T) {
	testCases := []TestCase{
		{
			name:     "simple",
			input:    []int{2, 4, 6},
			expected: 12,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Sum(tc.input)
			if result != tc.expected {
				t.Errorf("Ожидалось %d, получили %d", tc.expected, result)
			}
		})
	}
}
