package sumarray

func Sum(arr []int) int {
	sum := 0
	if len(arr) == 0 {
		return sum
	}

	return arr[0] + Sum(arr[1:])
}
