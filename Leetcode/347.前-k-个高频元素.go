/*
 * @Author: ronlee
 * @Date: 2021-08-10 23:07:05
 * @LastEditors: ronlee
 * @LastEditTime: 2021-08-15 11:42:49
 * @Description: file content
 * @FilePath: \Leetcode\347.前-k-个高频元素.go
 */
/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 *
 * https://leetcode-cn.com/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (62.19%)
 * Likes:    826
 * Dislikes: 0
 * Total Accepted:    185.2K
 * Total Submissions: 297.7K
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,1,1,2,2,3], k = 2
 * 输出: [1,2]
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [1], k = 1
 * 输出: [1]
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * k 的取值范围是 [1, 数组中不相同的元素的个数]
 * 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
 *
 *
 *
 *
 * 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
 *
 */

// @lc code=start
//前k高频--大顶堆

func topKFrequent(nums []int, k int) []int {
	//先统计频率
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	//大顶堆的优先队列，全部压入
	q := PriorityQueue{}
	for key, count := range m {
		heap.Push(&q, &Item{key: key, count: count})
	}
	var result []int
	//返回前k个
	for {
		if k <= 0 {
			break
		}
		result = append(result, heap.Pop(&q).(*Item).key)
		k--
	}
	return result
}

// Item define
type Item struct {
	key   int
	count int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	// 注意：因为 golang 中的 heap 默认是按最小堆组织的，所以 count 越大，Less() 越小，越靠近堆顶。这里采用 >，变为最大堆
	return (*pq)[i].count > (*pq)[j].count
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

// Push define
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

// Pop define
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

// @lc code=end

