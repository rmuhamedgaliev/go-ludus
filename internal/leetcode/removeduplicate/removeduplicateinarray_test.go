package removeduplicate

import (
	"fmt"
	"testing"
)

type TestCase struct {
	name        string
	array       []int
	expected    []int
	expectedNum int
}

func TestRemoveDuplicates(t *testing.T) {
	testCases := []TestCase{
		{
			name:        "simple",
			array:       []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected:    []int{0, 1, 2, 3, 4, -1, -1, -1, -1, -1},
			expectedNum: 5,
		},
		{
			name:        "simple",
			array:       []int{1, 1, 2},
			expected:    []int{1, 2, 2},
			expectedNum: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RemoveDuplicates(tc.array)
			if result != tc.expectedNum {
				t.Errorf("Ошибка ожидалось, что будет замен %d, а вышло %d", tc.expectedNum, result)
			}

			for i := 0; i < result; i++ {
				fmt.Printf("item: %d\n", tc.array[i])
				if tc.array[i] != tc.expected[i] {
					t.Errorf("Ошибка несоотвествия значения массива %d с индексом %d", tc.expected[i], i)
				}
			}
		})
	}
}
