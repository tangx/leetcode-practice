package duplicate04

import (
	"fmt"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	arr := [][]int{
		{1, 3, 4, 5},
		{2, 2, 4, 5},
	}

	for _, nums := range arr {
		ok := containsDuplicate(nums)
		fmt.Println(ok)
	}
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, val := range nums {
		if _, ok := m[val]; ok {
			return true
		}
		m[val] = false
	}

	return false
}
