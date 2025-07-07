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

	flag := true
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			flag = false
			break
		}
	}

	if flag {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}

}
