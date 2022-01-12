/*
 * @Author: ronlee
 * @Date: 2022-01-12 22:18:32
 * @LastEditors: ronlee
 * @LastEditTime: 2022-01-12 22:45:05
 * @Description: file content
 * @FilePath: \_1_learn\Note\Leetcode\303.区域和检索-数组不可变.go
 */
/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 *
 * https://leetcode-cn.com/problems/range-sum-query-immutable/description/
 *
 * algorithms
 * Easy (73.29%)
 * Likes:    406
 * Dislikes: 0
 * Total Accepted:    139.8K
 * Total Submissions: 190.7K
 * Testcase Example:  '["NumArray","sumRange","sumRange","sumRange"]\n' +
  '[[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]'
 *
 * 给定一个整数数组  nums，求出数组从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点。
 *
 *
 *
 * 实现 NumArray 类：
 *
 *
 * NumArray(int[] nums) 使用数组 nums 初始化对象
 * int sumRange(int i, int j) 返回数组 nums 从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点（也就是
 * sum(nums[i], nums[i + 1], ... , nums[j])）
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * ["NumArray", "sumRange", "sumRange", "sumRange"]
 * [[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
 * 输出：
 * [null, 1, -1, -3]
 *
 * 解释：
 * NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
 * numArray.sumRange(0, 2); // return 1 ((-2) + 0 + 3)
 * numArray.sumRange(2, 5); // return -1 (3 + (-5) + 2 + (-1))
 * numArray.sumRange(0, 5); // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * -10^5
 * 0
 * 最多调用 10^4 次 sumRange 方法
 *
 *
 *
 *
*/

// @lc code=start
type NumArray struct {
	Nums   []int
	Presum []int //用一个数组来存储前缀和
}

func Constructor(nums []int) NumArray {
	x := NumArray{
		Nums:   nums,
		Presum: make([]int, len(nums)), //初始化一个长度为nums的数组,不然无法索引取值
	}
	x.Presum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		x.Presum[i] = x.Presum[i-1] + nums[i]
	}
	return x
}

func (this *NumArray) SumRange(left int, right int) int {
	//这里注意相减后左边界被减没了
	return this.Presum[right] - this.Presum[left] + this.Nums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end

