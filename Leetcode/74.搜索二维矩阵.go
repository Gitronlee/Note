/*
 * @Author: ronlee
 * @Date: 2022-01-19 16:47:32
 * @LastEditors: ronlee
 * @LastEditTime: 2022-01-19 17:02:08
 * @Description: file content
 * @FilePath: \_1_learn\Note\Leetcode\74.搜索二维矩阵.go
 */
/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 *
 * https://leetcode-cn.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (46.67%)
 * Likes:    561
 * Dislikes: 0
 * Total Accepted:    194.9K
 * Total Submissions: 417.4K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3'
 *
 * 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
 *
 *
 * 每行中的整数从左到右按升序排列。
 * 每行的第一个整数大于前一行的最后一个整数。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1
 * -10^4
 *
 *
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	i := len(matrix) - 1
	if i < 0 {
		return false
	}
	j := len(matrix[0]) - 1
	if j < 0 {
		return false
	}
	// 从左下角开始查找，若目标值更小则往上走，更大则往右走
	curx := i
	cury := 0
	for curx >= 0 && cury <= j {
		if matrix[curx][cury] == target {
			return true
		} else if matrix[curx][cury] > target {
			curx--
		} else {
			cury++
		}
	}
	return false
}

// @lc code=end

