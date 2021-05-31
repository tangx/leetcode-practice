package twosum

func twoSum_01(nums []int, target int) []int {
	// 双循环 , 时间复杂度为 O(n^2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
