import "container/heap"

/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 *
 * https://leetcode-cn.com/problems/find-median-from-data-stream/description/
 *
 * algorithms
 * Hard (52.35%)
 * Likes:    635
 * Dislikes: 0
 * Total Accepted:    74K
 * Total Submissions: 141.2K
 * Testcase Example:  '["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]\n' +
  '[[],[1],[2],[],[3],[]]'
 *
 * 中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
 *
 * 例如，
 *
 * [2,3,4] 的中位数是 3
 *
 * [2,3] 的中位数是 (2 + 3) / 2 = 2.5
 *
 * 设计一个支持以下两种操作的数据结构：
 *
 *
 * void addNum(int num) - 从数据流中添加一个整数到数据结构中。
 * double findMedian() - 返回目前所有元素的中位数。
 *
 *
 * 示例：
 *
 * addNum(1)
 * addNum(2)
 * findMedian() -> 1.5
 * addNum(3)
 * findMedian() -> 2
 *
 * 进阶:
 *
 *
 * 如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
 * 如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
 *
 *
*/

// @lc code=start
type MedianFinder struct {
	// 用小顶堆保存较大的值，用大顶堆保存较小的值，则堆顶的中位数即为中位数
	large *IntLessHeap
	small *IntBigHead
}
type IntLessHeap []int
type IntBigHead []int

// 大小顶堆都要要实现5个接口方法其中Push Pop需要用指针且参数为interface{}
func (h IntLessHeap) Len() int {
	return len(h)
}
func (h IntBigHead) Len() int {
	return len(h)
}
// less确定堆为大顶还是小顶
func (h IntLessHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntBigHead) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h IntLessHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h IntBigHead) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IntLessHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntBigHead) Push(x interface{}) {
	*h = append(*h, x.(int))

}
func (h *IntLessHeap) Pop() interface{} {
	old := *h
	len := len(old)
	x := old[len-1]
	*h = (*h)[0 : len-1]
	return x
}
func (h *IntBigHead) Pop() interface{} {
	old := *h
	len := len(old)
	x := old[len-1]
	*h = (*h)[0 : len-1]
	return x

}
func Constructor() MedianFinder {
	a := &IntLessHeap{}
	b := &IntBigHead{}
	return MedianFinder{a, b}
}

func (this *MedianFinder) AddNum(num int) {
	if this.large.Len() >= this.small.Len() {
		//放入小顶
		heap.Push(this.large, num)
		heap.Push(this.small, heap.Pop(this.large).(int))
	} else {
		heap.Push(this.small, num)
		heap.Push(this.large, heap.Pop(this.small).(int))
	}
}
// 注意堆其实是int切片，堆顶就是[0]，注意用引用方式获取切片在取其下标值
func (this *MedianFinder) FindMedian() float64 {
	if this.large.Len() == this.small.Len() {
		return ((float64((*this.small)[0]) + float64((*this.large)[0])) / 2.0)
	} else {
		return float64((*this.small)[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end
