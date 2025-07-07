package main

import (
	"fmt"
)

func main() {
	var countOfNumbers int
	_, _ = fmt.Scan(&countOfNumbers)

	arr := make([]int, countOfNumbers)

	for index := range arr {
		_, _ = fmt.Scan(&arr[index])
	}

	seen := make(map[int]bool)
	flag := false

	for _, number := range arr {
		if seen[number] {
			flag = true
			break
		}
		seen[number] = true
	}

	if flag {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
