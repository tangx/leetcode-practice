package removeduplicates

// removeDuplicates20200530 删除重复值
// 解题思路， 使用双指针
func removeDuplicates20200530(nums []int) int {

	// 如果长度为 0
	// nil 长度也为0
	if len(nums) == 0 {
		return 0
	}

	// 如果长度不为0
	head := 0

	// 移动指针 tail
	for tail := 1; tail < len(nums); tail++ {
		// 如果 nums[tail] 与 nums[head] 相同， 则 tail 继续移动，
		if nums[tail] == nums[head] {
			continue
		}

		// 如果 nums[tail] 与 nums[head] 不相同， 则 head+1 并复制 nums[tail]
		head += 1
		nums[head] = nums[tail]
	}

	// 返回长度 head+1
	// 因为 head 指针位置从 0 开始， 因此长度要 +1
	return head + 1
}
