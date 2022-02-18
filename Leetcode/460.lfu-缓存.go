import "container/list"

/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU 缓存
 *
 * https://leetcode-cn.com/problems/lfu-cache/description/
 *
 * algorithms
 * Hard (43.52%)
 * Likes:    494
 * Dislikes: 0
 * Total Accepted:    39.4K
 * Total Submissions: 90.5K
 * Testcase Example:  '["LFUCache","put","put","get","put","get","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]'
 *
 * 请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。
 *
 * 实现 LFUCache 类：
 *
 *
 * LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
 * int get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1 。
 * void put(int key, int value) - 如果键 key 已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量
 * capacity 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用
 * 的键。
 *
 *
 * 为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。
 *
 * 当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。
 *
 * 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get",
 * "get"]
 * [[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
 * 输出：
 * [null, null, null, 1, null, -1, 3, null, -1, 3, 4]
 *
 * 解释：
 * // cnt(x) = 键 x 的使用计数
 * // cache=[] 将显示最后一次使用的顺序（最左边的元素是最近的）
 * LFUCache lfu = new LFUCache(2);
 * lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
 * lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
 * lfu.get(1);      // 返回 1
 * ⁠                // cache=[1,2], cnt(2)=1, cnt(1)=2
 * lfu.put(3, 3);   // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
 * ⁠                // cache=[3,1], cnt(3)=1, cnt(1)=2
 * lfu.get(2);      // 返回 -1（未找到）
 * lfu.get(3);      // 返回 3
 * ⁠                // cache=[3,1], cnt(3)=2, cnt(1)=2
 * lfu.put(4, 4);   // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
 * ⁠                // cache=[4,3], cnt(4)=1, cnt(3)=2
 * lfu.get(1);      // 返回 -1（未找到）
 * lfu.get(3);      // 返回 3
 * ⁠                // cache=[3,4], cnt(4)=1, cnt(3)=3
 * lfu.get(4);      // 返回 4
 * ⁠                // cache=[3,4], cnt(4)=2, cnt(3)=3
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= capacity <= 10^4
 * 0 <= key <= 10^5
 * 0 <= value <= 10^9
 * 最多调用 2 * 10^5 次 get 和 put 方法
 *
 *
*/

// @lc code=start
type LFUCache struct {
	nodes    map[int]*list.Element
	lists    map[int]*list.List
	capacity int
	min      int
}

type node struct {
	key       int
	value     int
	frequency int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{nodes: make(map[int]*list.Element),
		lists:    make(map[int]*list.List),
		capacity: capacity,
		min:      0,
	}
}

// 向列表中添加元素
func (this *LFUCache) Add2List(key int, currentNode *node) {
	// 检查对应的列表是否存在不存在则创建
	if _, ok := this.lists[currentNode.frequency]; !ok {
		this.lists[currentNode.frequency] = list.New()
	}
	// 将元素添加到列表中，并更新map
	newList := this.lists[currentNode.frequency]
	newNode := newList.PushBack(currentNode)
	this.nodes[key] = newNode
}

// 更新最小频率的值
func (this *LFUCache) UpdateMin(currentNode *node) {
	// 如果新添加到队列的值频率为1，则min为1
	if currentNode.frequency == 1 {
		this.min = 1
		return
	}
	// 如果之前的频率是最小频率，且因更新导致其队列删空，则min++
	if currentNode.frequency-1 == this.min && this.lists[currentNode.frequency-1].Len() == 0 {
		this.min++
		return
	}
}

// 被使用一次则更新其频率并维护新的队列
func (this *LFUCache) UseOnece(key int) {
	value, _ := this.nodes[key]
	currentNode := value.Value.(*node)
	// 删除其在旧队列中的元素
	this.lists[currentNode.frequency].Remove(value)
	// 更新其频率
	currentNode.frequency++
	// 往新队列中添加
	this.Add2List(key, currentNode)
	// 更新最小频率
	this.UpdateMin(currentNode)
}

// 获取值
func (this *LFUCache) Get(key int) int {
	// 不存在直接返回-1
	value, ok := this.nodes[key]
	if !ok {
		return -1
	}
	// 存在则断言为node并取值返回，并调用UseOnece
	currentNode := value.Value.(*node)
	this.UseOnece(key)
	return currentNode.value
}

func (this *LFUCache) Put(key int, value int) {
	// 容量为0直接返回
	if this.capacity == 0 {
		return
	}
	// 如果key存在，则更新value
	if currentValue, ok := this.nodes[key]; ok {
		// 1、节点内的值做更新
		currentNode := currentValue.Value.(*node)
		currentNode.value = value
		// 2、节点被用了一次
		this.UseOnece(key)
		return
	}
	// 如果满了，则删除最小频率的节点
	if this.capacity == len(this.nodes) {
		// 获取最小频率节点的队列
		currentList := this.lists[this.min]
		frontNode := currentList.Front()
		// 删除map中删除节点kv对，从此队列中删除节点
		delete(this.nodes, frontNode.Value.(*node).key)
		currentList.Remove(frontNode)
	}
	// 创建新节点并加入队列，更新min
	currentNode := &node{
		key:       key,
		value:     value,
		frequency: 1,
	}
	this.Add2List(key, currentNode)
	this.UpdateMin(currentNode)

}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
