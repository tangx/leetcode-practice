/*
旋转图像
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。



示例 1：


输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[[7,4,1],[8,5,2],[9,6,3]]
示例 2：


输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
示例 3：

输入：matrix = [[1]]
输出：[[1]]
示例 4：

输入：matrix = [[1,2],[3,4]]
输出：[[3,1],[4,2]]


提示：

matrix.length == n
matrix[i].length == n
1 <= n <= 20
-1000 <= matrix[i][j] <= 1000
相关标签
数组

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnhhkv/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/
package rotatematrix

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

// rotate 旋转矩阵
// 类似的， 数据元素循序旋转， 通常都可以使用 **反转** 来解决问题
// 例如， 03_ratate_array
// 解题思路
// 1. 先 matrix 进行 / 对角线翻转
// 2. 在对翻转后的 - 水平线垂直翻转
func rotate(matrix [][]int) {

	// 这里需要注意边界为题， 因为索引是从 0 开始
	// ( row = length-1-col )
	// 3=(4-1)-1
	// 2=3-2
	// 1=3-3
	// 0=3-4
	edge := len(matrix) - 1

	// 1. 第一次对角线交换
	for row := 0; row <= edge; row++ {
		for col := 0; col <= edge; col++ {

			// 翻转，直到遇到中线, 进行下一次外层循环
			if row+col > edge {
				break
			}

			// 否则进行 cell 值交换
			// row=0,col=1 : (0,1)<=>(3,2)=(3-0,3-1)=(edge-row,edge-col)
			// row=0,col=2 : (0,2)<=>(3,1)=(edge-0,edge-2)=(edge-row,edge-col)
			// row=2,col=0 : (2,0)<=>(1,3)=(edge-2,edge-0)=(edge-row,edge-col)
			// row=1,col=1 : (1,1)<=>(2,2)=(edge-1,edge-1)=(edge-row,edge-col)
			matrix[row][col], matrix[edge-col][edge-row] = matrix[edge-col][edge-row], matrix[row][col]
			fmt.Println(matrix)
		}
	}

	// todo // 2. 水平线交换
	for col := 0; col <= edge; col++ {
		for row := 0; row <= edge-row; row++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}
}

func Test_RotateMatrix(t *testing.T) {
	datas := []struct {
		Result [][]int
		Data   [][]int
	}{
		{
			Result: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
			Data:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		}, {
			Result: [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			Data:   [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
	}

	for _, data := range datas {
		rotate(data.Data)

		gomega.NewWithT(t).Expect(data.Data).To(gomega.Equal(data.Result))
	}
}
