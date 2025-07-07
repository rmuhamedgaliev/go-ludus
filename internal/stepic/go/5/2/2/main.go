package main

import "fmt"

func main() {
	var countOfNumbers int
	_, _ = fmt.Scan(&countOfNumbers)

	arr := make([]int, countOfNumbers)

	minIndex := 0
	for index := range arr {
		_, _ = fmt.Scan(&arr[index])
		if arr[minIndex] > arr[index] {
			minIndex = index
		}
	}

	minValue := arr[minIndex]
	for index := range arr {
		arr[index] = arr[index] - minValue
	}

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
}
