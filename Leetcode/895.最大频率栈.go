import "container/heap"

/*
 * @Author: ronlee
 * @Date: 2022-02-18 16:20:02
 * @LastEditors: ronlee
 * @LastEditTime: 2022-02-18 17:09:08
 * @Description: file content
 * @FilePath: \_1_learn\Note\Leetcode\895.最大频率栈.go
 */
/*
 * @lc app=leetcode.cn id=895 lang=golang
 *
 * [895] 最大频率栈
 *
 * https://leetcode-cn.com/problems/maximum-frequency-stack/description/
 *
 * algorithms
 * Hard (56.65%)
 * Likes:    207
 * Dislikes: 0
 * Total Accepted:    9.1K
 * Total Submissions: 16.1K
 * Testcase Example:  '["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"]\n' +
  '[[],[5],[7],[5],[7],[4],[5],[],[],[],[]]'
 *
 * 设计一个类似堆栈的数据结构，将元素推入堆栈，并从堆栈中弹出出现频率最高的元素。
 *
 * 实现 FreqStack 类:
 *
 *
 * FreqStack() 构造一个空的堆栈。
 * void push(int val) 将一个整数 val 压入栈顶。
 * int pop() 删除并返回堆栈中出现频率最高的元素。
 *
 * 如果出现频率最高的元素不只一个，则移除并返回最接近栈顶的元素。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：
 *
 * ["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"],
 * [[],[5],[7],[5],[7],[4],[5],[],[],[],[]]
 * 输出：[null,null,null,null,null,null,null,5,7,5,4]
 * 解释：
 * FreqStack = new FreqStack();
 * freqStack.push (5);//堆栈为 [5]
 * freqStack.push (7);//堆栈是 [5,7]
 * freqStack.push (5);//堆栈是 [5,7,5]
 * freqStack.push (7);//堆栈是 [5,7,5,7]
 * freqStack.push (4);//堆栈是 [5,7,5,7,4]
 * freqStack.push (5);//堆栈是 [5,7,5,7,4,5]
 * freqStack.pop ();//返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,5,7,4]。
 * freqStack.pop ();//返回 7 ，因为 5 和 7 出现频率最高，但7最接近顶部。堆栈变成 [5,7,5,4]。
 * freqStack.pop ();//返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,4]。
 * freqStack.pop ();//返回 4 ，因为 4, 5 和 7 出现频率最高，但 4 是最接近顶部的。堆栈变成 [5,7]。
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= val <= 10^9
 * push 和 pop 的操作数不大于 2 * 10^4。
 * 输入保证在调用 pop 之前堆栈中至少有一个元素。
 *
 *
*/

// @lc code=start
type FreqStack struct {
	// 使用一个小顶堆，每次插入一个元素，都把元素和其出现次数放入小顶堆中，
	// 每次弹出一个元素，都把元素和其出现次数从小顶堆中弹出，
	items ItemHeap
	// 记录对应key值的频率
	freq map[int]int
	//当前时间戳
	time int
}
type Item struct {
	val  int // 元素值
	freq int // 出现频率
	age  int // 元素的插入顺序
}
type ItemHeap []Item

// 小顶堆需要实现5个接口：Len、Less、Swap、Push、Pop
func (s ItemHeap) Len() int {
	return len(s)
}
func (s ItemHeap) Less(i, j int) bool {
	if s[i].freq == s[j].freq {
		return s[i].age > s[j].age //频率相等时，返回后入的
	}
	return s[i].freq > s[j].freq //频率不相等时，返回频率大的
}
func (s ItemHeap) Swap(i, j int) {
	// 交换需要用到指针
	s[i], s[j] = s[j], s[i]
}
func (s *ItemHeap) Push(x interface{}) {
	*s = append(*s, x.(Item))
}
func (s *ItemHeap) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}
func Constructor() FreqStack {
	return FreqStack{
		items: make(ItemHeap, 0),
		freq:  make(map[int]int),
		time:  0,
	}
}

func (this *FreqStack) Push(val int) {
	this.freq[val]++
	heap.Push(&this.items,
		Item{val: val, freq: this.freq[val], age: this.time})
	this.time++
}

func (this *FreqStack) Pop() int {
	// 弹出操作直接从小顶堆中弹出
	item := heap.Pop(&this.items).(Item)
	this.freq[item.val]--

	return item.val
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
// @lc code=end
