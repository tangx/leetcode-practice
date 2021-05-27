/*
	买卖股票的最佳时机 II
	给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。

	设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

	注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。



	示例 1:

	输入: prices = [7,1,5,3,6,4]
	输出: 7
	解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
	     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
	示例 2:

	输入: prices = [1,2,3,4,5]
	输出: 4
	解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
	     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
	示例 3:

	输入: prices = [7,6,4,3,1]
	输出: 0
	解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。


	提示：

	1 <= prices.length <= 3 * 104
	0 <= prices[i] <= 104
	相关标签
	贪心算法
	数组

	作者：力扣 (LeetCode)
	链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2zsx1/
	来源：力扣（LeetCode）
	著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

// maxProfit 最大盈利
// 贪婪算法: 最多次的完成交易， 只要后一天比前一天有赚， 就在前一天买，后一天卖。
func maxProfit(prices []int) int {
	// 边界处理
	// 如果为 nil, 0-1 天数据， 则不必购买
	if len(prices) < 1 {
		return 0
	}

	// 正常逻辑
	sum := 0
	for i := 1; i < len(prices); i++ {
		if earn := prices[i] - prices[i-1]; earn > 0 {
			sum += earn
		}
	}

	return sum
}

func Test_maxProfit(t *testing.T) {

	// 边界处理
	t.Run("nil, or 0-1 days", func(t *testing.T) {
		NewWithT(t).Expect(maxProfit(nil)).To(Equal(0))
		NewWithT(t).Expect(maxProfit([]int{})).To(Equal(0))
		NewWithT(t).Expect(maxProfit([]int{1})).To(Equal(0))
	})

	// 盈利
	t.Run("few days", func(t *testing.T) {

		data := map[int][]int{
			7: {7, 1, 5, 3, 6, 4},
			4: {1, 2, 3, 4, 5},
			0: {7, 6, 4, 3, 1},
		}
		for sum, prices := range data {
			NewWithT(t).Expect(maxProfit(prices)).To(Equal(sum))
		}
	})
}
