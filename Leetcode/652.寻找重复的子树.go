/*
 * @Author: ronlee
 * @Date: 2021-12-28 20:22:23
 * @LastEditors: ronlee
 * @LastEditTime: 2021-12-29 10:06:10
 * @Description: file content
 * @FilePath: \_1_learn\Note\Leetcode\652.寻找重复的子树.go
 */
/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树
 *
 * https://leetcode-cn.com/problems/find-duplicate-subtrees/description/
 *
 * algorithms
 * Medium (57.08%)
 * Likes:    351
 * Dislikes: 0
 * Total Accepted:    37.3K
 * Total Submissions: 65.2K
 * Testcase Example:  '[1,2,3,4,null,2,4,null,null,4]'
 *
 * 给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
 *
 * 两棵树重复是指它们具有相同的结构以及相同的结点值。
 *
 * 示例 1：
 *
 * ⁠       1
 * ⁠      / \
 * ⁠     2   3
 * ⁠    /   / \
 * ⁠   4   2   4
 * ⁠      /
 * ⁠     4
 *
 *
 * 下面是两个重复的子树：
 *
 * ⁠     2
 * ⁠    /
 * ⁠   4
 *
 *
 * 和
 *
 * ⁠   4
 *
 *
 * 因此，你需要以列表的形式返回上述重复子树的根结点。
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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	hashAll := map[string]int{}
	res := []*TreeNode{}
	dfs(root, hashAll, &res)
	return res
}

func dfs(node *TreeNode, hashAll map[string]int, res *[]*TreeNode) string {
	if node == nil {
		return "#"
	}
	lString := dfs(node.Left, hashAll, res)
	rString := dfs(node.Right, hashAll, res)
	//序列化时主语这里的括号，若不加括号可能导致【00#】和【0#0】的误判
	buildString := fmt.Sprintf("(%s)(%d)(%s)", lString, node.Val, rString)
	hashAll[buildString]++
	if hashAll[buildString] == 2 {
		*res = append(*res, node)
	}
	return buildString
}

// @lc code=end

