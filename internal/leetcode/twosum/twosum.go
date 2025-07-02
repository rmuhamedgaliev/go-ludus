package twosum

func TwoSum(nums []int, target int) []int {
	seenNumbers := make(map[int]int, len(nums))
	result := []int{-1, -1}
	for index, value := range nums {
		required := target - value

		mapValue, ok := seenNumbers[required]
		if ok {
			result[0] = mapValue
			result[1] = index
			break
		} else {
			seenNumbers[value] = index
		}
	}

	return result
}
