package haspathsum

import (
	"testing"

	"github.com/onsi/gomega"
)

var GenMockDatas = []struct {
	summary string
	data    []interface{}
}{
	// {summary: "nil", data: nil},
	// {summary: "empty", data: []interface{}{}},
	{summary: "len=1", data: []interface{}{1}},
	{summary: "len=2", data: []interface{}{1, nil}},
	{summary: "len > 2 with nil", data: []interface{}{1, 2, 3, nil, 5, 6, 7}},
}

func Test_Generator(t *testing.T) {
	for _, mock := range GenMockDatas {
		t.Run(mock.summary, func(t *testing.T) {
			_, err := generator(mock.data)

			gomega.NewWithT(t).Expect(err).To(gomega.BeNil())
		})
	}
}
