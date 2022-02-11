/*
 * @lc app=leetcode.cn id=1373 lang=golang
 *
 * [1373] 二叉搜索子树的最大键值和
 *
 * https://leetcode-cn.com/problems/maximum-sum-bst-in-binary-tree/description/
 *
 * algorithms
 * Hard (40.64%)
 * Likes:    74
 * Dislikes: 0
 * Total Accepted:    9.4K
 * Total Submissions: 23.2K
 * Testcase Example:  '[1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]'
 *
 * 给你一棵以 root 为根的 二叉树 ，请你返回 任意 二叉搜索子树的最大键值和。
 *
 * 二叉搜索树的定义如下：
 *
 *
 * 任意节点的左子树中的键值都 小于 此节点的键值。
 * 任意节点的右子树中的键值都 大于 此节点的键值。
 * 任意节点的左子树和右子树都是二叉搜索树。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root = [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]
 * 输出：20
 * 解释：键值为 3 的子树是和最大的二叉搜索树。
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：root = [4,3,null,1,2]
 * 输出：2
 * 解释：键值为 2 的单节点子树是和最大的二叉搜索树。
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [-4,-2,-5]
 * 输出：0
 * 解释：所有节点键值都为负数，和最大的二叉搜索树为空。
 *
 *
 * 示例 4：
 *
 *
 * 输入：root = [2,1,3]
 * 输出：6
 *
 *
 * 示例 5：
 *
 *
 * 输入：root = [5,4,8,3,null,6,3]
 * 输出：7
 *
 *
 *
 *
 * 提示：
 *
 *
 * 每棵树有 1 到 40000 个节点。
 * 每个节点的键值在 [-4 * 10^4 , 4 * 10^4] 之间。
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
func maxSumBST(root *TreeNode) int {
	res := 0
	dfs(root, &res)
	return res
}
func dfs(root *TreeNode, res *int) []int {
	if root == nil {
		//分别为 是否为BST，当前树中的最小值，最大值，最大键值和
		// (这里的初始最大最小值注意选，最小值的边界为max，最大值的边界为min，以对应min和max函数做更新)
		return []int{1, 40001, -40001, 0}
	}
	left := dfs(root.Left, res)
	right := dfs(root.Right, res)
	// 若左右子树都是BST，且root值大于左子树最大值，小于右子树最小值，则root为BST，则返回1，最小值，最大值，最大键值和
	if left[0] == 1 && right[0] == 1 && root.Val > left[2] && root.Val < right[1] {
		//所以当前root为BST，且用当前键值和尝试更新最大键值和
		*res = max(*res, left[3]+right[3]+root.Val)
		return []int{1, min(left[1], root.Val), max(right[2], root.Val), left[3] + right[3] + root.Val}
	} else {
		return []int{0, 0, 0, 0}
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

