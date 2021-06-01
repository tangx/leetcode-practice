/*
请你判断一个 9x9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
数独部分空格内已填入了数字，空白格用 '.' 表示。

注意：

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。


示例 1：


输入：board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
输出：true
示例 2：

输入：board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
输出：false
解释：除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。 但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。


提示：

board.length == 9
board[i].length == 9
board[i][j] 是一位数字或者 '.'
相关标签
哈希表

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2f9gg/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
package isvalidsudoku

import (
	"testing"

	"github.com/onsi/gomega"
)

// isValidSudoku
// 使用 hashset 表， 以格子中的 value 作为 hashset 的 key。
// 并判断 hashset key 是否存在
func isValidSudoku(board [][]byte) bool {

	// 这是一个二维数组
	// 要求
	// board[1][y] : 第一行不能有重复
	row := make(map[byte]struct{})

	// board[x][1] : 第一列不能有重复
	col := make(map[byte]struct{})

	// board[x/3][y/3] : 第 n 块不能有重复。 1/3=0..1, 4/3=1..1, 7/3=2..1
	// cell 除了行之外， 还有列。 因此 cell 的表达式应该为 x/3 + (y/3)*3
	// 0 1 2
	// 3 4 5
	// 6 7 8
	cell := make(map[byte]struct{})

	// 双循环
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {

			val := board[x][y]

			// 如果格子为点， 则跳过
			if val == '.' {
				continue
			}

			// 如果格子为值， 则判断

			// 因为无法初始化 [9]map[byte]struct{},
			// 所以不得不以 9*x + val 的结果作为 set 的key
			r := byte(x) * '9'
			if _, ok := row[r+val]; ok {
				return false
			}

			c := byte(y) * '9'
			if _, ok := col[c+val]; ok {
				return false
			}

			k := byte(x/3+y/3*3) * '9'
			if _, ok := cell[k+val]; ok {
				return false
			}

			row[r+val] = struct{}{}
			col[c+val] = struct{}{}
			cell[k+val] = struct{}{}
		}
	}

	return true
}

func Test_IsValidSudoku(t *testing.T) {
	board := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	actual := isValidSudoku(board)
	gomega.NewWithT(t).Expect(actual).To(gomega.BeFalse())

	board2 := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	actual2 := isValidSudoku(board2)
	gomega.NewWithT(t).Expect(actual2).To(gomega.BeTrue())
}