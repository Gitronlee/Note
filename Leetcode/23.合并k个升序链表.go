/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 *
 * https://leetcode-cn.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (55.65%)
 * Likes:    1438
 * Dislikes: 0
 * Total Accepted:    296.2K
 * Total Submissions: 531.8K
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 给你一个链表数组，每个链表都已经按升序排列。
 *
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 *
 *
 *
 * 示例 1：
 *
 * 输入：lists = [[1,4,5],[1,3,4],[2,6]]
 * 输出：[1,1,2,3,4,4,5,6]
 * 解释：链表数组如下：
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * 将它们合并到一个有序链表中得到。
 * 1->1->2->3->4->4->5->6
 *
 *
 * 示例 2：
 *
 * 输入：lists = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 * 输入：lists = [[]]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] 按 升序 排列
 * lists[i].length 的总和不超过 10^4
 *
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
type node_slice_less_heap []*ListNode

func (n node_slice_less_heap) Len() int {
	return len(n)
}
func (n node_slice_less_heap) Less(i, j int) bool {
	return n[i].Val < n[j].Val
}
func (n *node_slice_less_heap) Swap(i, j int) {
	(*n)[i], (*n)[j] = (*n)[j], (*n)[i]
}
func (n *node_slice_less_heap) Push(h interface{}) {
	node := h.(*ListNode)
	*n = append(*n, node)
}
func (n *node_slice_less_heap) Pop() interface{} {
	temp := (*n)[len(*n)-1]
	*n = (*n)[:len(*n)-1]
	return temp
}
func mergeKLists(lists []*ListNode) *ListNode {
	//go 的heap包支持构建大（小）顶堆
	//此处构建一个小顶堆，需要实现对应接口方法Len Less（<即小顶堆）
	//Swap Pop Push 需要传参为指针，不然没有结果返回
	h := node_slice_less_heap{}
	//基于head去调用堆
	heap.Init(&h)
	for _, v := range lists {
		if v != nil {
			heap.Push(&h, v)
		}
	}
	head := &ListNode{}
	cur := head
	for h.Len() > 0 {
		//逻辑上为将K个链表一次入堆，再一次pop最小，但若此链还有值则后续再入堆、
		ln := heap.Pop(&h).(*ListNode)
		cur.Next = ln
		if ln.Next != nil {
			heap.Push(&h, ln.Next)
		}
		cur = cur.Next
	}
	return head.Next
}

// @lc code=end

