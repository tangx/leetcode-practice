package isvalidsudoku

// isValidSudoku_02
// 解题思路: seti[i][val] || setj[j][val] || setk[k][val]
// 使用一个二维数组, 保存行/列/块中是否出现在该数字， 如果没有， 则设置为 true
// seti[1][6]  : 第一行是否出现过 6 这个值。
func isValidSudoku_02(board [][]byte) bool {

	// cell[i][j] // 一纬

	/*
		seti[0][5]=true // seti[i][val]=true
		0: 第 0 行
		5: 单元格数字 5
		true: 存在 / false: 不存在

		第「0」行有数字「5」存在。
	*/

	// 数独的大小是固定的， 所以用数组而非切片
	// 行
	seti := [9][9]bool{}
	// 列
	setj := [9][9]bool{}

	/*
		将大九宫81格分成9个小九宫9格
		0 1 2
		3 4 5
		6 7 8

		setk[k][val]=true
		计算方式:
		k:=i/3+(j/3)*3
			i=4,j=0 -> 1+0*3=1
			i=0,j=4 -> 0+1*3=3
	*/
	// 小九宫格
	setk := [9][9]bool{}

	for i := 0; i < len(board); i++ {

		// 二纬
		for j := 0; j < len(board[i]); j++ {

			// 点就略过
			if board[i][j] == '.' {
				continue
			}

			/*
				val := board[i][j]

				因为原值是 byte 类型， 因此超出了 1-9 的范围。
				panic: runtime error: index out of range [56] with length 9 [recovered]
				panic: runtime error: index out of range [56] with length 9

				1. 「 - '0' 」 确认 1-9 的范围
				2. 「 - 1 」 数组各自从 0 开始
			*/
			val := board[i][j] - '0' - 1

			k := i/3 + j/3*3

			// 如果 行， 列， 小九宫 中有 val 值， 则为不合法
			if seti[i][val] || setj[j][val] || setk[k][val] {
				return false
			}

			// val 不层出现， 设置 flag
			seti[i][val], setj[j][val], setk[k][val] = true, true, true
		}

	}
	return true
}
