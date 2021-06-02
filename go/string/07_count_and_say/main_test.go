package countandsay

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

var MockDatas = []struct {
	n int
	s string
}{
	// {n: 0, s: ""},
	{n: 1, s: "1"},
	{n: 2, s: "11"},
	{n: 3, s: "21"},
	{n: 4, s: "1211"},
	{n: 5, s: "111221"},
}

func Test_CountString(t *testing.T) {
	s := countString("11122211")
	gomega.NewWithT(t).Expect(s).To(gomega.Equal("313221"))
}

func Test_CountAndSay(t *testing.T) {
	for _, mock := range MockDatas {
		t.Run(fmt.Sprint(mock), func(t *testing.T) {
			s := countAndSay(mock.n)
			gomega.NewWithT(t).Expect(s).To(gomega.Equal(mock.s))
		})
	}
}
