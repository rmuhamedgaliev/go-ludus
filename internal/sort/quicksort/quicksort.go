package quicksort

func Sort(arr []int, left int, right int) {
	if left >= right {
		return
	}

	pivotIndex := partition(arr, left, right)

	Sort(arr, left, pivotIndex-1)
	Sort(arr, pivotIndex+1, right)
}

func partition(arr []int, left int, right int) int {
	pivot := arr[right]
	i := left - 1

	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}
