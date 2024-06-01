package Leetcode

func twoSum(nums []int, target int) []int {
	var numbers []int
	for i := 0; i < len(nums); i++ {

		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target && j != i {
				numbers = append(numbers, i, j)
				return numbers
			}
		}
	}
	return numbers
}
