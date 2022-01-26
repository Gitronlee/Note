/*
 * @Author: ronlee
 * @Date: 2022-01-26 10:37:32
 * @LastEditors: ronlee
 * @LastEditTime: 2022-01-26 22:39:56
 * @Description: file content
 * @FilePath: \_1_learn\Note\Leetcode\19.删除链表的倒数第-n-个结点.go
 */
/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 *
 * https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (43.35%)
 * Likes:    1769
 * Dislikes: 0
 * Total Accepted:    628.9K
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], n = 2
 * 输出：[1,2,3,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1], n = 1
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = [1,2], n = 1
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中结点的数目为 sz
 * 1 <= sz <= 30
 * 0 <= Node.val <= 100
 * 1 <= n <= sz
 *
 *
 *
 *
 * 进阶：你能尝试使用一趟扫描实现吗？
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{-1, head}
	fast, slow := dummyHead, dummyHead
	// fast和slow的距离为n
	for i := 0; i < n; i++ {
		if fast == nil {
			return head
		}
		fast = fast.Next
	}
	//fast到达最后一个节点时，slow到达倒数第n个节点
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// dummy的第n个节点其实是list的第n-1个节点，倒数第n个节点是list的倒数第n+1个节点
	// 所以删除slow的下个节点
	slow.Next = slow.Next.Next
	return dummyHead.Next
}

// @lc code=end

