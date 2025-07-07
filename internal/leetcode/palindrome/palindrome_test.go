package palindrome

import "testing"

type TestCase struct {
	name     string
	number   int
	expected bool
}

func TestIsPalindrome(t *testing.T) {
	testCases := []TestCase{
		{
			name:     "simple",
			number:   121,
			expected: true,
		},
		{
			name:     "negative",
			number:   -121,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsPalindrome(tc.number)
			if result != tc.expected {
				t.Errorf("Ошибка, ожидалось %t, вышло %t", tc.expected, result)
			}
		})
	}
}
