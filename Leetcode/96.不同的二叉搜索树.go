/*
 * @Author: ronlee
 * @Date: 2022-01-05 15:47:28
 * @LastEditors: ronlee
 * @LastEditTime: 2022-02-10 15:11:36
 * @Description: file content
 * @FilePath: \_1_learn\Note\Leetcode\96.不同的二叉搜索树.go
 */
/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees/description/
 *
 * algorithms
 * Medium (69.92%)
 * Likes:    1472
 * Dislikes: 0
 * Total Accepted:    187.5K
 * Total Submissions: 267.9K
 * Testcase Example:  '3'
 *
 * 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：5
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：1
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
func numTrees(n int) int {
	//从1到n的个数
	c := make(map[string]int, 0)
	return count(1, n, c)
}
func count(lo int, hi int, c map[string]int) int {
	//用一个map记录是否已经计算过
	cur := fmt.Sprintf("%d-%d", lo, hi)
	if val, ok := c[cur]; ok {
		return val
	}
	// 不合法区间时只能生成1个nil的树
	if lo > hi {
		c[cur] = 1
		return 1
	}
	res := 0
	//当前的[lo,hi]的数字可组成的bst中由各各数为根节点时左*右的累加得到。
	for i := lo; i <= hi; i++ {
		left := count(lo, i-1, c)
		right := count(i+1, hi, c)
		res += left * right
	}
	c[cur] = res
	return res
}

// @lc code=end

