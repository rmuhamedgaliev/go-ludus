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

	fmt.Print(minIndex)
}
