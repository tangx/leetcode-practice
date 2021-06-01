package isanagram

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

var MockDatas = []struct {
	s  string
	t  string
	ok bool
}{
	{s: "anagram", t: "nagaram", ok: true},
	{s: "rat", t: "car", ok: false},
	{s: "a", t: "ab", ok: false},
	{s: "aaaaarrrrr", t: "raaraarrra", ok: true},
	{s: "abcdefg", t: "defghki", ok: false},
}

func Test_IsAnagram(t *testing.T) {

	f := func(fn func(s string, t string) bool) {
		for _, mock := range MockDatas {

			t.Run(fmt.Sprint(mock), func(t *testing.T) {
				ok := fn(mock.s, mock.t)
				gomega.NewWithT(t).Expect(ok).To(gomega.Equal(mock.ok))
			})
		}
	}

	f(isAnagram__sort)
	f(isAnagram__loop)
	f(isAnagram__loopOnce)
}
