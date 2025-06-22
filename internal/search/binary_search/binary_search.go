package binarysearch

func Search(arr []int, number int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)>>1

		if arr[mid] == number {
			return mid
		}

		if arr[mid] > number {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
