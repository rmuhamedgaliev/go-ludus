package selectionsort

func Sort(arr []int) []int {
	for index := range arr {
		smallestIndex := index
		for j := index + 1; j < len(arr); j++ {
			if arr[j] < arr[smallestIndex] {
				smallestIndex = j
			}
		}
		arr[index], arr[smallestIndex] = arr[smallestIndex], arr[index]
	}
	return arr
}
