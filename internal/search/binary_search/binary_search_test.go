package binarysearch

import "testing"

type TestCase struct {
	name     string
	arr      []int
	number   int
	expected int
}

func TestSearch(t *testing.T) {
	testCases := []TestCase{
		{
			name:     "simple",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			number:   3,
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Search(tc.arr, tc.number)

			if result != tc.expected {
				t.Errorf("Ожидалось %d, а получили %d", tc.expected, result)
			}
		})
	}
}

func BenchmarkLarge(b *testing.B) {
	arr := make([]int, 1000000) 
	for i := 0; i < 1000000; i++ {
		arr[i] = i * 2
	}
	number := 500000
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Search(arr, number)
	}
}
