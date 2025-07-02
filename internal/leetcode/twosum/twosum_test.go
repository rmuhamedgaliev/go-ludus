package twosum

import (
	"testing"
)

type TestCase struct {
	name     string
	arr      []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	testCases := []TestCase{
		{
			name:     "simple",
			arr:      []int{1, 2, 3},
			target:   4,
			expected: []int{0, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := TwoSum(tc.arr, tc.target)
			if tc.expected[0] != result[0] || tc.expected[1] != result[1] {
				t.Errorf("Ошибка ожидали %d и %d, получили %d и %d", tc.expected[0], tc.expected[1], result[0], result[1])
			}
		})
	}
}
