/*
二叉树的前序遍历
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。



示例 1：


输入：root = [1,null,2,3]
输出：[1,2,3]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]
示例 4：


输入：root = [1,2]
输出：[1,2]
示例 5：


输入：root = [1,null,2]
输出：[1,2]


提示：

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100


进阶：递归算法很简单，你可以通过迭代算法完成吗？

相关标签
栈
树

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/data-structure-binary-tree/xeywh5/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
package preordertraversal

import (
	"fmt"
	"testing"
)

// Test_Slice 参数传递
func Test_Slice(t *testing.T) {
	sl := []int{}
	fmt.Printf("start: sl ptr= %p \n", sl)
	reserve := func(i int, sl []int) {
		sl = append(sl, i)
		fmt.Printf("  sl(%d) ptr= %p => ", i, sl)

		fmt.Println(sl)
	}
	for i := 0; i < 5; i++ {
		reserve(i, sl)
	}
	fmt.Printf("end  : sl ptr= %p \n", sl)

	// === RUN   Test_Slice
	// start: sl ptr= 0x124af98
	//   sl(0) ptr= 0xc0000a01b0 => [0]
	//   sl(1) ptr= 0xc0000a01c0 => [1]
	//   sl(2) ptr= 0xc0000a01d0 => [2]
	//   sl(3) ptr= 0xc0000a01e0 => [3]
	//   sl(4) ptr= 0xc0000a01f0 => [4]
	// end  : sl ptr= 0x124af98
	// --- PASS: Test_Slice (0.00s)
}

// Test_Slice02_Closure 闭包
func Test_Slice02_Closure(t *testing.T) {
	sl := []int{}
	fmt.Printf("start: sl ptr= %p \n", sl)
	reserve := func(i int) {
		sl = append(sl, i)
		fmt.Printf("  sl(%d) ptr= %p => ", i, sl)

		fmt.Println(sl)
	}
	for i := 0; i < 5; i++ {
		reserve(i)
	}
	fmt.Printf("end  : sl ptr= %p ->%x\n", sl, sl)

	// === RUN   Test_Slice02_Closure
	// start: sl ptr= 0x124bfb8
	//   sl(0) ptr= 0xc0000a0220 => [0]
	//   sl(1) ptr= 0xc0000a0230 => [0 1]
	//   sl(2) ptr= 0xc0000d6080 => [0 1 2]
	//   sl(3) ptr= 0xc0000d6080 => [0 1 2 3]
	//   sl(4) ptr= 0xc0000dc040 => [0 1 2 3 4]
	// end  : sl ptr= 0xc0000dc040 ->[0 1 2 3 4]
	// --- PASS: Test_Slice02_Closure (0.00s)
}

func Test_SliceAppendNil(t *testing.T) {
	sl := []int{}
	sl2 := func() []int {
		return nil
	}()
	sl = append(sl, sl2...)
	fmt.Println(sl)

	// === RUN   Test_SliceAppendNil
	// []
	// --- PASS: Test_SliceAppendNil (0.00s)
	// PASS
}
