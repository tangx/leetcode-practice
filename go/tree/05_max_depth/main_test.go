package maxdepth

import (
	"testing"

	"github.com/onsi/gomega"
)

func Test_MaxDepth(t *testing.T) {
	r := maxDepth(MockNode)
	gomega.NewWithT(t).Expect(r).To(gomega.Equal(3))
}
