/*
 * @Author: ronlee
 * @Date: 2021-08-15 22:58:03
 * @LastEditors: ronlee
 * @LastEditTime: 2021-08-15 23:30:00
 * @Description: file content
 * @FilePath: \Leetcode\264.丑数-ii.go
 */
/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 *
 * https://leetcode-cn.com/problems/ugly-number-ii/description/
 *
 * algorithms
 * Medium (57.41%)
 * Likes:    707
 * Dislikes: 0
 * Total Accepted:    90.4K
 * Total Submissions: 157.3K
 * Testcase Example:  '10'
 *
 * 给你一个整数 n ，请你找出并返回第 n 个 丑数 。
 *
 * 丑数 就是只包含质因数 2、3 和/或 5 的正整数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 10
 * 输出：12
 * 解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：1
 * 解释：1 通常被视为丑数。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 */

// @lc code=start
func nthUglyNumber(n int) int {
	//容易想到当前的丑数肯定来自前边某个丑数乘2或3或5得到的。
	//那么我们维护一个待选数组【two，three， five】并记录他们当前的坐标i2,i3，i5.任意一个被选走后，只要对应ix++乘，即得新的替补
	if n == 1 {
		return 1
	}
	//最初的替补选手就是235
	two, three, five := 2, 3, 5
	i2, i3, i5 := 0, 0, 0
	res := make([]int, n)
	res[0] = 1
	for i := 1; i < n; i++ {
		//若当前替补选手有重复，搞成不重复
		if two == three || two == five {
			i2++
			two = res[i2] * 2
		}
		if three == five {
			i3++
			three = res[i3] * 3
		}
		if two < three && two < five {
			res[i] = two
			i2++
			two = res[i2] * 2
		} else if three < two && three < five {
			res[i] = three
			i3++
			three = res[i3] * 3
		} else if five < three && five < two {
			res[i] = five
			i5++
			five = res[i5] * 5
		}
	}
	return res[n-1]
}

// @lc code=end

