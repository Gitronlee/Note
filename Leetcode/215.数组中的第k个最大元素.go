/*
 * @Author: ronlee
 * @Date: 2021-08-09 22:19:10
 * @LastEditors: ronlee
 * @LastEditTime: 2021-08-09 23:00:41
 * @Description: file content
 * @FilePath: \Leetcode\215.数组中的第k个最大元素.go
 */
/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (64.70%)
 * Likes:    1213
 * Dislikes: 0
 * Total Accepted:    396.2K
 * Total Submissions: 612.2K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
 *
 * 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 *
 *
 * 示例 2:
 *
 *
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10^4
 *
 *
 */

// @lc code=start
// 方法一，小顶堆法
// type lessheap []int

// func (h lessheap) Len() int { return len(h) }
// func (h lessheap) Less(i, j int) bool {
// 	return h[i] < h[j]
// }
// func (h lessheap) Swap(i, j int) {
// 	h[i], h[j] = h[j], h[i]
// }
// func (h lessheap) Peek() int {
// 	return h[0]
// }
// func (h *lessheap) Pop() interface{} {
// 	n := len(*h)
// 	res := (*h)[n-1]
// 	*h = (*h)[:n-1]
// 	return res
// }
// func (h *lessheap) Push(x interface{}) {
// 	*h = append((*h), x.(int))
// }
// func findKthLargest(nums []int, k int) int {
// 	//最优解时基于快排的分治思想，此处先实现一个基于小顶堆的实现
// 	//小顶堆k个元素，当前值大于堆顶则压入，并弹出顶。最终就维持了k个最大值。而此时的堆顶就是第k大值（堆中最小值）
// 	if k <= 0 || len(nums) <= 0 {
// 		return 0
// 	}
// 	h := &lessheap{}
// 	heap.Init(h)
// 	for i := 0; i < k; i++ {
// 		heap.Push(h, nums[i])
// 	}
// 	for i := k; i < len(nums); i++ {
// 		if nums[i] > h.Peek() {
// 			heap.Push(h, nums[i])
// 			heap.Pop(h)
// 		}
// 	}
// 	return h.Peek()
// }
// 方法二，快排分治
func findKthLargest(nums []int, k int) int {
	//快排思想把基准数大的都放在基准数的右边,把比基准数小的放在基准数的左边
	if k <= 0 || len(nums) <= 0 {
		return 0
	}
	//首先定义函数，再[lf=eft, right]直接返回left处的值排序后的下标
	partition := func(nums []int, left, right int) int {
		pivot := nums[left] //临时记录
		j := left
		for i := left + 1; i <= right; i++ {
			if nums[i] < pivot {
				j++ //记录排序后pivot的下标
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
		//归位
		nums[j], nums[left] = nums[left], nums[j]
		return j
	}
	n := len(nums)
	left, right, target := 0, n-1, n-k
	for {
		//若恰好返回的就是第n-k个 也就是第k大，直接返回；否则缩小左右边界
		index := partition(nums, left, right)
		if index == target {
			return nums[index]
		} else if index < target {
			left = index + 1
		} else {
			right = index - 1
		}
	}
}

// @lc code=end

