/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
 *
 * https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/description/
 *
 * algorithms
 * Medium (75.22%)
 * Likes:    546
 * Dislikes: 0
 * Total Accepted:    174.8K
 * Total Submissions: 232.3K
 * Testcase Example:  '[3,1,4,null,2]\n1'
 *
 * 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,1,4,null,2], k = 1
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [5,3,6,2,4,null,null,1], k = 3
 * 输出：3
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数为 n 。
 * 1
 * 0
 *
 *
 *
 *
 * 进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？
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
func kthSmallest(root *TreeNode, k int) int {
	//第k个,即中序遍历的第k个，中序遍历时对c赋值，并依次取完前k-1个，则最后一个即结果
	c := make(chan int)
	go inorderfunc(root, c)
	for i := 1; i <= k-1; i++ {
		<-c
	}
	return <-c
}
func inorderfunc(root *TreeNode, c chan int) {
	if root == nil {
		return
	}
	inorderfunc(root.Left, c)
	c <- root.Val
	inorderfunc(root.Right, c)
}

// @lc code=end

