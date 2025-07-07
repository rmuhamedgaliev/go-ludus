package main

import "fmt"

func main() {
	var countOfNumbers int
	_, _ = fmt.Scan(&countOfNumbers)

	arr := make([]int, countOfNumbers)

	minIndex := 0
	maxIndex := 0
	for index := range arr {
		_, _ = fmt.Scan(&arr[index])
		if arr[minIndex] > arr[index] {
			minIndex = index
		}

		if arr[index] >= arr[maxIndex] {
			maxIndex = index
		}
	}

	tmp := arr[minIndex]
	arr[minIndex] = arr[maxIndex]
	arr[maxIndex] = tmp

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
}
