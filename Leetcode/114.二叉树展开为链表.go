/*
 * @Author: ronlee
 * @Date: 2021-12-01 15:58:47
 * @LastEditors: ronlee
 * @LastEditTime: 2021-12-01 20:53:34
 * @Description: file content
 * @FilePath: \_1_learn\Note\Leetcode\114.二叉树展开为链表.go
 */
/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 *
 * https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/description/
 *
 * algorithms
 * Medium (72.65%)
 * Likes:    985
 * Dislikes: 0
 * Total Accepted:    193.3K
 * Total Submissions: 266K
 * Testcase Example:  '[1,2,5,3,4,null,6]'
 *
 * 给你二叉树的根结点 root ，请你将它展开为一个单链表：
 *
 *
 * 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
 * 展开后的单链表应该与二叉树 先序遍历 顺序相同。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,5,3,4,null,6]
 * 输出：[1,null,2,null,3,null,4,null,5,null,6]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [0]
 * 输出：[0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中结点数在范围 [0, 2000] 内
 * -100
 *
 *
 *
 *
 * 进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？
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
func flatten(root *TreeNode) {
	//递归思路：
	//1.如果当前节点为空，则直接返回
	if root == nil {
		return
	}
	//2.将左右子树展开为单链表
	flatten(root.Left)
	flatten(root.Right)
	//3.将左子树拼在root的右树，并向右遍历到底，再将之前保存的右子树拼接上
	temp := root.Right
	root.Right = root.Left
	root.Left = nil
	toright := root
	for {
		if toright.Right == nil {
			break
		}
		toright = toright.Right
	}
	toright.Right = temp
}

// @lc code=end

