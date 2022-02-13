/*
 * @Author: ronlee
 * @Date: 2022-02-11 11:24:42
 * @LastEditors: ronlee
 * @LastEditTime: 2022-02-11 14:46:29
 * @Description: file content
 * @FilePath: \_1_learn\Note\Leetcode\146.lru-缓存.go
 */
/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 *
 * https://leetcode-cn.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (52.49%)
 * Likes:    1902
 * Dislikes: 0
 * Total Accepted:    279.9K
 * Total Submissions: 533.3K
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
 * 
 * 实现 LRUCache 类：
 * 
 * 
 * 
 * 
 * LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
 * int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
 * void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
 * key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
 * 
 * 
 * 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
 * 
 * 
 * 
 * 
 * 
 * 示例：
 * 
 * 
 * 输入
 * ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
 * [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
 * 输出
 * [null, null, null, 1, null, -1, null, -1, 3, 4]
 * 
 * 解释
 * LRUCache lRUCache = new LRUCache(2);
 * lRUCache.put(1, 1); // 缓存是 {1=1}
 * lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
 * lRUCache.get(1);    // 返回 1
 * lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
 * lRUCache.get(2);    // 返回 -1 (未找到)
 * lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
 * lRUCache.get(1);    // 返回 -1 (未找到)
 * lRUCache.get(3);    // 返回 3
 * lRUCache.get(4);    // 返回 4
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= capacity <= 3000
 * 0 <= key <= 10000
 * 0 <= value <= 10^5
 * 最多调用 2 * 10^5 次 get 和 put
 * 
 * 
 */

// @lc code=start
type LRUCache struct {
	// LRUCache为最近最少使用缓存策略，淘汰近期最少使用的值。因此用双向链表来实现。
	// 双向链表可以在指导curNode时对链表在O1时间复杂度进行操作。
	// 由于get和put需要在O1，故需要map来得到key到Node的映射。
	// Cap用来记录容量
	head, tail *Node
	Keys       map[int]*Node
	Cap        int
}
// 双向链表要维护前驱和后继，这里的Val用来计算Cache的值，Key值用来记录其在map中的key。便于根据链表节点反找map的key-v
type Node struct {
	Key, Val   int
	Prev, Next *Node
}
// 初始化时头尾节点为nil
func Constructor(capacity int) LRUCache {
	return LRUCache{Keys: make(map[int]*Node), Cap: capacity, head: nil, tail: nil}
}
// get操作时从map中取值，并将其移动到链表头部，这里需要注意的是，如果key不存在，则返回-1。
func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.Val
	}
	return -1
}
// put操作也要看是否已存在，若存在则更新其值，并用remove+add移动到链表头部
func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		node.Val = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		// 不存在直接将新节点添加到链表头部，增加key-v映射
		node = &Node{Key: key, Val: value}
		this.Keys[key] = node
		this.Add(node)
	}
	// 若超过最大容量应删除最后一个节点和其对应的key-v映射
	if len(this.Keys) > this.Cap {
		delete(this.Keys, this.tail.Key)
		this.Remove(this.tail)
	}
}
// 向链表头部增加节点
func (this *LRUCache) Add(node *Node) {
	// 增加后新的节点就是头节点，因此其前驱为nil，后继为原头节点
	node.Prev = nil
	node.Next = this.head
	// 若原来的头结点不为nil那么原来的头结点变成普通节点其prev指向新节点
	if this.head != nil {
		this.head.Prev = node
	}
	// 标记新的头结点
	this.head = node
	// 添加时同时判断当前尾节点是否为空，若空则新节点也被标记为尾节点
	if this.tail == nil {
		this.tail = node
		this.tail.Next = nil
	}
}
// 移除节点
func (this *LRUCache) Remove(node *Node) {
// 判断要删除的节点是否是头结点
	if node == this.head {
		this.head = node.Next
		node.Next = nil
		return
	}
	// 也可能同事是尾节点
	if node == this.tail {
		this.tail = node.Prev
		node.Prev.Next = nil
		node.Prev = nil
		return
	}
	// 若未return则是删除中间的节点
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

