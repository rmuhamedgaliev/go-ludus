package selectionsort

func Sort(arr []int) []int {
	result := make([]int, 0, len(arr))
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)

	for range arr {
		smallest := FindSmallestIndex(arrCopy)
		result = append(result, arrCopy[smallest])
		arrCopy = append(arrCopy[:smallest], arrCopy[smallest+1:]...)
	}
	return result
}

func FindSmallestIndex(arr []int) int {
	smallest_index := 0
	smallest := arr[smallest_index]

	for index, item := range arr {
		if item < smallest {
			smallest = item
			smallest_index = index
		}
	}

	return smallest_index
}
