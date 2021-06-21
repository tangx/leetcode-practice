package ispalindromechain

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

func Test_ReverseChain(t *testing.T) {
	list := []int{1, 2, 3, 4}
	chain := NewChain(list)

	ptr := chain
	for ptr != nil {
		fmt.Println(ptr.Val)
		ptr = ptr.Next
	}

	chain2 := reverse(chain)

	ptr = chain2
	for ptr != nil {
		fmt.Println(ptr.Val)
		ptr = ptr.Next
	}
}

func Test_IsPalindrome(t *testing.T) {

	mockDatas := []struct {
		data   []int
		result bool
	}{
		{
			data:   []int{1},
			result: true,
		},
		{
			data:   []int{1, 2, 3, 4},
			result: false,
		},
		{
			data:   []int{1, 2, 2, 1},
			result: true,
		},
		{
			data:   []int{1, 2, 3, 2, 1},
			result: true,
		},
	}

	for _, mock := range mockDatas {

		t.Run(fmt.Sprint(mock), func(t *testing.T) {
			head := NewChain(mock.data)
			ret := isPalindrome(head)

			gomega.NewWithT(t).Expect(ret).To(gomega.Equal(mock.result))
		})
	}
}
