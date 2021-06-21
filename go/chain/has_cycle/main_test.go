package hascycle

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

func Test_hasCycle(t *testing.T) {

	mockDatas := []struct {
		data   []int
		pos    int
		result bool
	}{
		{[]int{3, 2, 0, -4}, 1, true},
		{[]int{1, 2}, 0, true},
		{[]int{1}, -1, false},
		{[]int{1}, 0, true},
		{[]int{1, 2}, -1, false},
	}

	for _, mock := range mockDatas {
		t.Run(fmt.Sprint(mock), func(t *testing.T) {

			chain := NewChain(mock.data, mock.pos)
			ok := hasCycle(chain)

			gomega.NewWithT(t).Expect(ok).To(gomega.Equal(mock.result))
		})
	}
}
