package quicksort

import "testing"

func TestSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}

	Sort(arr, 0, len(arr)-1)

	for i, v := range arr {
		if v != expected[i] {
			t.Errorf("Deifferent at index %d: got %d, want %d", i, v, expected[i])
		}
	}

}
