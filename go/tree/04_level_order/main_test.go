package levelorder

import (
	"fmt"
	"testing"
)

func Test_LevelOrder(t *testing.T) {
	answer := levelOrder(MockNode)
	fmt.Println(answer)
}
