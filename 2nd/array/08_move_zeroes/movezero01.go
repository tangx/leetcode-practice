package movezeroes

func moveZeroes01(nums []int) {
	// 需要移动 0 到的数组末尾位置
	head := 0
	tail := 0

	// head, tail 两个指针同时从左向右移动
	// 如果 head 为0， 则表示该位置需替换。
	// 如果 tail 不为0， 则表示该位置需要为 head 换位。
	for tail < len(nums) {
		switch {
		case nums[head] == 0 && nums[tail] != 0:
			nums[head], nums[tail] = nums[tail], nums[head]
			head++
			tail++
		case nums[head] != 0:
			head++
			fallthrough
		case nums[tail] == 0:
			tail++
		}
	}
}
