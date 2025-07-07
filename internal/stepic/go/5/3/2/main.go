package main

import "fmt"

func main() {
	numbers := []int{1, 1, 1, 2, 2}

	cnt := 0
	numbersSize := len(numbers)
	for i, number := range numbers {
		for j := i + 1; j < numbersSize; j++ {
			if number == numbers[j] {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
