package threesum

import "testing"

type TestCase struct {
	name     string
	nums     []int
	target   int
	expected []int
}

func TestThreeSum(t *testing.T) {
	testCases := []TestCase{
		{
			name:     "simple",
			nums:     []int{1, 2, 3, 4, 5, 6},
			target:   7,
			expected: []int{0, 1, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ThreeSum(tc.nums, tc.target)
			for i := 0; i < 3; i++ {
				if tc.expected[i] != result[i] {
					t.Errorf("Ошибка ожидали %d, а получили %d", tc.expected[i], result[i])
				}
			}
		})
	}
}
