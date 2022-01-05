/*
 * @Author: ronlee
 * @Date: 2022-01-05 21:04:54
 * @LastEditors: ronlee
 * @LastEditTime: 2022-01-05 21:37:40
 * @Description: file content
 * @FilePath: \_1_learn\Note\Leetcode\95.不同的二叉搜索树-ii.go
 */
/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (70.81%)
 * Likes:    1088
 * Dislikes: 0
 * Total Accepted:    110.9K
 * Total Submissions: 156.7K
 * Testcase Example:  '3'
 *
 * 给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：[[1]]
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
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return build(1, n)
}
func build(lo int, hi int) []*TreeNode {
	//用于记录结果
	var res []*TreeNode
	if lo > hi {
		res = append(res, nil)
		return res
	}
	for i := lo; i <= hi; i++ {
		//i为根节点时，先构建出左右子树的所有情况
		leftList := build(lo, i-1)
		rightList := build(i+1, hi)
		for li := 0; li < len(leftList); li++ {
			for ri := 0; ri < len(rightList); ri++ {
				// 将其两两组合
				cur := TreeNode{
					Val:   i,
					Left:  leftList[li],
					Right: rightList[ri]} //注意go的这里不能另起一行
				res = append(res, &cur)
			}
		}
	}
	return res
}

// @lc code=end

