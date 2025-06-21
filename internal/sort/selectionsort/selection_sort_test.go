package selectionsort

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"
)

type TestCase struct {
	name     string
	arr      []int
	expected []int
}

func TestSort(t *testing.T) {
	testCases := []TestCase{
		{
			name:     "empty array",
			arr:      []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			arr:      []int{42},
			expected: []int{42},
		},
		{
			name:     "already sorted",
			arr:      []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse sorted",
			arr:      []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "random order",
			arr:      []int{7, 2, 9, 1, 5},
			expected: []int{1, 2, 5, 7, 9},
		},
		{
			name:     "duplicates",
			arr:      []int{3, 1, 3, 2, 1},
			expected: []int{1, 1, 2, 3, 3},
		},
		{
			name:     "all same elements",
			arr:      []int{5, 5, 5, 5},
			expected: []int{5, 5, 5, 5},
		},
		{
			name:     "negative numbers",
			arr:      []int{-3, -1, -4, -2},
			expected: []int{-4, -3, -2, -1},
		},
		{
			name:     "mixed positive and negative",
			arr:      []int{3, -1, 0, -2, 5},
			expected: []int{-2, -1, 0, 3, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			originalCopy := make([]int, len(tc.arr))
			copy(originalCopy, tc.arr)

			result := Sort(tc.arr)

			if !slices.Equal(result, tc.expected) {
				t.Errorf("Sort(%v) = %v, ожидали %v", originalCopy, result, tc.expected)
			}

			if !slices.Equal(tc.arr, originalCopy) {
				t.Errorf("Оригинальный массив был изменен: был %v, стал %v", originalCopy, tc.arr)
			}
		})
	}
}

func TestSortLarge(t *testing.T) {
	sizes := []int{100, 500, 1000}

	for _, size := range sizes {
		t.Run(fmt.Sprintf("size_%d", size), func(t *testing.T) {
			rng := rand.New(rand.NewSource(42))
			arr := make([]int, size)
			for i := range arr {
				arr[i] = rng.Intn(1000) - 500
			}

			originalCopy := make([]int, len(arr))
			copy(originalCopy, arr)

			result := Sort(arr)

			if !slices.IsSorted(result) {
				t.Error("Результат не отсортирован")
			}

			if len(result) != len(originalCopy) {
				t.Errorf("Длина изменилась: была %d, стала %d", len(originalCopy), len(result))
			}

			slices.Sort(originalCopy)
			if !slices.Equal(result, originalCopy) {
				t.Error("Элементы массива изменились или потерялись")
			}
		})
	}
}

func TestSortDoesNotModifyInput(t *testing.T) {
	original := []int{5, 2, 8, 1, 9}
	backup := make([]int, len(original))
	copy(backup, original)

	Sort(original)

	if !slices.Equal(original, backup) {
		t.Errorf("Входной массив был изменен: был %v, стал %v", backup, original)
	}
}

func BenchmarkSort_Random_100(b *testing.B) {
	benchmarkSort(b, generateRandom(100))
}

func BenchmarkSort_Random_500(b *testing.B) {
	benchmarkSort(b, generateRandom(500))
}

func BenchmarkSort_Random_1000(b *testing.B) {
	benchmarkSort(b, generateRandom(1000))
}

func BenchmarkSort_Sorted_1000(b *testing.B) {
	benchmarkSort(b, generateSorted(1000))
}

func BenchmarkSort_Reverse_1000(b *testing.B) {
	benchmarkSort(b, generateReverse(1000))
}

func BenchmarkSort_Duplicates_1000(b *testing.B) {
	benchmarkSort(b, generateDuplicates(1000))
}

func benchmarkSort(b *testing.B, original []int) {
	testCopy := make([]int, len(original))
	copy(testCopy, original)
	testResult := Sort(testCopy)

	if len(testResult) != len(original) {
		b.Fatalf("Неверная длина результата: получили %d, ожидали %d", len(testResult), len(original))
	}

	if len(testResult) > 1 && !slices.IsSorted(testResult) {
		b.Fatal("Функция возвращает неотсортированный результат")
	}

	originalSorted := make([]int, len(original))
	copy(originalSorted, original)
	slices.Sort(originalSorted)
	if !slices.Equal(testResult, originalSorted) {
		b.Fatal("Функция теряет или изменяет элементы")
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		arr := make([]int, len(original))
		copy(arr, original)
		Sort(arr)
	}
}

func generateRandom(size int) []int {
	rng := rand.New(rand.NewSource(42))
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rng.Intn(2000) - 1000
	}
	return arr
}

func generateSorted(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i
	}
	return arr
}

func generateReverse(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = size - i - 1
	}
	return arr
}

func generateDuplicates(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i % 10
	}
	return arr
}
