package threesum

import (
	"rmuhamedgaliev.dev/go-ludus/internal/leetcode/twosum"
)

func ThreeSum(nums []int, target int) []int {
	result := make([]int, 3)

	for i := 0; i < len(nums)-2; i++ {
		first := nums[i]
		targetForTwo := target - first
		cutIndex := i + 1
		twoResult := twosum.TwoSum(nums[cutIndex:], targetForTwo)
		if twoResult[0] > -1 && twoResult[1] > -1 {
			result[0] = i
			result[1] = twoResult[0] + cutIndex
			result[2] = twoResult[1] + cutIndex
			break
		}
	}

	return result
}
