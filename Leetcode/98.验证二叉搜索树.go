/*
 * @Author: ronlee
 * @Date: 2022-01-06 11:07:35
 * @LastEditors: ronlee
 * @LastEditTime: 2022-01-06 14:10:42
 * @Description: file content
 * @FilePath: \Leetcode\98.验证二叉搜索树.go
 */
/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode-cn.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (35.30%)
 * Likes:    1370
 * Dislikes: 0
 * Total Accepted:    398.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '[2,1,3]'
 *
 * 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
 *
 * 有效 二叉搜索树定义如下：
 *
 *
 * 节点的左子树只包含 小于 当前节点的数。
 * 节点的右子树只包含 大于 当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [2,1,3]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [5,1,4,null,null,3,6]
 * 输出：false
 * 解释：根节点的值是 5 ，但是右子节点的值是 4 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目范围在[1, 10^4] 内
 * -2^31 <= Node.val <= 2^31 - 1
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
func isValidBST(root *TreeNode) bool {
	return RecValidate(root, nil, nil)
}

func RecValidate(n, min, max *TreeNode) bool {
	// 如果当前为空，返回true
	if n == nil {
		return true
	}
	// 如果当前不为空，则当前的值必须在给定的范围内，否则返回false
	if min != nil && n.Val <= min.Val {
		return false
	}
	if max != nil && n.Val >= max.Val {
		return false
	}
	//若当前值在给定范围内，则其左右支树也必须在给定范围内：对左右子树这个范围可以相应的缩小
	return RecValidate(n.Left, min, n) && RecValidate(n.Right, n, max)
}

// @lc code=end

