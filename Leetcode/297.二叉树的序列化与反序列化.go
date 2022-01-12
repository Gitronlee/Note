/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 *
 * https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/description/
 *
 * algorithms
 * Hard (56.43%)
 * Likes:    723
 * Dislikes: 0
 * Total Accepted:    119.8K
 * Total Submissions: 212.3K
 * Testcase Example:  '[1,2,3,null,null,4,5]'
 *
 *
 * 序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
 *
 * 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 /
 * 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
 *
 * 提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode
 * 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,3,null,null,4,5]
 * 输出：[1,2,3,null,null,4,5]
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
 * 输入：root = [1]
 * 输出：[1]
 *
 *
 * 示例 4：
 *
 *
 * 输入：root = [1,2]
 * 输出：[1,2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中结点数在范围 [0, 10^4] 内
 * -1000
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

type Codec struct {
	builder strings.Builder // 用来保存二叉树序列化的字符串
	input   []string        // 用来保存字符串列表用于反序列化为二叉树
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	//序列化
	if root == nil {
		this.builder.WriteString("#,")
		return ""
	}
	//先当前
	this.builder.WriteString(strconv.Itoa(root.Val) + ",")
	//再左
	this.serialize(root.Left)
	//再右
	this.serialize(root.Right)
	// 全部序列化完成后返回
	return this.builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	// 反序列化的输入是一个字符串，先将其转为字符串数组
	this.input = strings.Split(data, ",")
	// 所以input里的元素为#或 数字，用deserializeNode函数来解析
	return this.deserializeHelper()
}
func (this *Codec) deserializeHelper() *TreeNode {
	// 先中间的元素,是否为空 # 并删除第一个元素
	if this.input[0] == "#" {
		this.input = this.input[1:]
		return nil
	}
	// 如果不是#，则说明是数字，则取出来
	val, _ := strconv.Atoi(this.input[0])
	// 将当前元素从input中删除
	this.input = this.input[1:]
	// 创建当前节点
	root := &TreeNode{Val: val}
	// 左节点
	root.Left = this.deserializeHelper()
	// 右节点
	root.Right = this.deserializeHelper()
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

