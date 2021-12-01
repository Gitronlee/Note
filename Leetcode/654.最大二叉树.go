/*
 * @Author: ronlee
 * @Date: 2021-12-01 21:53:20
 * @LastEditors: ronlee
 * @LastEditTime: 2021-12-01 22:04:23
 * @Description: file content
 * @FilePath: \_1_learn\Note\Leetcode\654.最大二叉树.go
 */
/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
 *
 * https://leetcode-cn.com/problems/maximum-binary-tree/description/
 *
 * algorithms
 * Medium (81.00%)
 * Likes:    347
 * Dislikes: 0
 * Total Accepted:    66.7K
 * Total Submissions: 82.4K
 * Testcase Example:  '[3,2,1,6,0,5]'
 *
 * 给定一个不含重复元素的整数数组 nums 。一个以此数组直接递归构建的 最大二叉树 定义如下：
 *
 *
 * 二叉树的根是数组 nums 中的最大元素。
 * 左子树是通过数组中 最大值左边部分 递归构造出的最大二叉树。
 * 右子树是通过数组中 最大值右边部分 递归构造出的最大二叉树。
 *
 *
 * 返回有给定数组 nums 构建的 最大二叉树 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [3,2,1,6,0,5]
 * 输出：[6,3,5,null,2,0,null,null,1]
 * 解释：递归调用如下所示：
 * - [3,2,1,6,0,5] 中的最大值是 6 ，左边部分是 [3,2,1] ，右边部分是 [0,5] 。
 * ⁠   - [3,2,1] 中的最大值是 3 ，左边部分是 [] ，右边部分是 [2,1] 。
 * ⁠       - 空数组，无子节点。
 * ⁠       - [2,1] 中的最大值是 2 ，左边部分是 [] ，右边部分是 [1] 。
 * ⁠           - 空数组，无子节点。
 * ⁠           - 只有一个元素，所以子节点是一个值为 1 的节点。
 * ⁠   - [0,5] 中的最大值是 5 ，左边部分是 [0] ，右边部分是 [] 。
 * ⁠       - 只有一个元素，所以子节点是一个值为 0 的节点。
 * ⁠       - 空数组，无子节点。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,2,1]
 * 输出：[3,null,2,null,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 * nums 中的所有整数 互不相同
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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	//定义一个函数用于获得 nums 中的最大值和索引
	var helper func(nums []int) (int, int)
	helper = func(nums []int) (int, int) {
		max := nums[0]
		maxIndex := 0
		for i, v := range nums {
			if v > max {
				max = v
				maxIndex = i
			}
		}
		return max, maxIndex
	}
	//若 nums 为空，则返回空
	if len(nums) == 0 {
		return nil
	}
	//否则将最大值给root，数组拆分后给左右子树
	max, maxIndex := helper(nums)
	root := &TreeNode{Val: max}
	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	return root
}

// @lc code=end

