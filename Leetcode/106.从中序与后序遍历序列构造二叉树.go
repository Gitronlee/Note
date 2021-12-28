/*
 * @Author: ronlee
 * @Date: 2021-12-28 09:42:15
 * @LastEditors: ronlee
 * @LastEditTime: 2021-12-28 09:45:55
 * @Description: file content
 * @FilePath: \_1_learn\Note\Leetcode\106.从中序与后序遍历序列构造二叉树.go
 */
/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (72.16%)
 * Likes:    642
 * Dislikes: 0
 * Total Accepted:    150K
 * Total Submissions: 207.8K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 根据一棵树的中序遍历与后序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 中序遍历 inorder = [9,3,15,20,7]
 * 后序遍历 postorder = [9,15,7,20,3]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	//与105类似 空直接返回
	if len(inorder) == 0 {
		return nil
	}
	//获取后序遍历的最后一个元素在中序中的位置
	idx := indexof(inorder, postorder[len(postorder)-1])
	//拆分为左右子树
	return &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  buildTree(inorder[:idx], postorder[:idx]),
		Right: buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1]),
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

