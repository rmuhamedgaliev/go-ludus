package removeduplicate

func RemoveDuplicates(nums []int) int {
	j := 0 // указатель для записи
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] { // сравниваем с последним уникальным
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
