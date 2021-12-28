/*
 * @Author: ronlee
 * @Date: 2021-12-28 09:05:53
 * @LastEditors: ronlee
 * @LastEditTime: 2021-12-28 09:15:07
 * @Description: file content
 * @FilePath: \_1_learn\Note\Leetcode\105.从前序与中序遍历序列构造二叉树.go
 */
/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (70.57%)
 * Likes:    1357
 * Dislikes: 0
 * Total Accepted:    279.1K
 * Total Submissions: 395.2K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 给定一棵树的前序遍历 preorder 与中序遍历  inorder。请构造二叉树并返回其根节点。
 *
 *
 *
 * 示例 1:
 *
 *
 * Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
 * Output: [3,9,20,null,null,15,7]
 *
 *
 * 示例 2:
 *
 *
 * Input: preorder = [-1], inorder = [-1]
 * Output: [-1]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1
 * inorder.length == preorder.length
 * -3000
 * preorder 和 inorder 均无重复元素
 * inorder 均出现在 preorder
 * preorder 保证为二叉树的前序遍历序列
 * inorder 保证为二叉树的中序遍历序列
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	//递归基 空时直接返回
	if len(preorder) == 0 {
		return nil
	}
	//否则拆分为左右子树
	idx := indexof(inorder, preorder[0])
	return &TreeNode{
		Val: preorder[0],
		//这里注意前闭后开区间
		Left:  buildTree(preorder[1:idx+1], inorder[:idx]),
		Right: buildTree(preorder[idx+1:], inorder[idx+1:]),
	}

}
func indexof(inorder []int, val int) int {
	for i, v := range inorder {
		if v == val {
			return i
		}
	}
	return -1
}

// @lc code=end

