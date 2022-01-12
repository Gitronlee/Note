/*
 * @Author: ronlee
 * @Date: 2022-01-12 21:43:45
 * @LastEditors: ronlee
 * @LastEditTime: 2022-01-12 22:01:53
 * @Description: file content
 * @FilePath: \_1_learn\Note\Leetcode\236.二叉树的最近公共祖先.go
 */
/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 *
 * https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/description/
 *
 * algorithms
 * Medium (68.40%)
 * Likes:    1479
 * Dislikes: 0
 * Total Accepted:    297.5K
 * Total Submissions: 434.9K
 * Testcase Example:  '[3,5,1,6,2,0,8,null,null,7,4]\n5\n1'
 *
 * 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
 *
 * 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x
 * 的深度尽可能大（一个节点也可以是它自己的祖先）。”
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
 * 输出：3
 * 解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
 * 输出：5
 * 解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1,2], p = 1, q = 2
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [2, 10^5] 内。
 * -10^9
 * 所有 Node.val 互不相同 。
 * p != q
 * p 和 q 均存在于给定的二叉树中。
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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		//如果当前为nil 或 当前为p 或当前为q 则直接返回当前
		return root
	}
	//否则从左右子树中查找
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	//如果分别在左右子树中都找到了
	if left != nil && right != nil {
		//说明当前为左右子树的最近公共祖先，有basecase这里两边都是非空则必是找到qp
		//若找到的是nil，那么就是没有找到qp，其公共祖先也就是nil
		return root
	}
	//如果都是空
	if left != nil {
		return left
	}

	return right
}

// @lc code=end

