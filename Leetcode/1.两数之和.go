/*
 * @Author: ronlee
 * @Date: 2021-08-02 21:09:15
 * @LastEditors: ronlee
 * @LastEditTime: 2021-08-02 21:15:37
 * @Description: file content
 * @FilePath: \Note\Leetcode\1.两数之和.go
 */
/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 *
 * https://leetcode-cn.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (51.69%)
 * Likes:    11716
 * Dislikes: 0
 * Total Accepted:    2.3M
 * Total Submissions: 4.5M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
 *
 * 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
 *
 * 你可以按任意顺序返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,7,11,15], target = 9
 * 输出：[0,1]
 * 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,2,4], target = 6
 * 输出：[1,2]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [3,3], target = 6
 * 输出：[0,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2
 * -10^9
 * -10^9
 * 只会存在一个有效答案
 *
 *
 * 进阶：你可以想出一个时间复杂度小于 O(n^2) 的算法吗？
 *
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	//用一个map记录是否another的值对应的索引被保存过
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		//若被保存过，则直接返回
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		//未返回时，将当前值与坐标保存
		m[nums[i]] = i
	}
	return []int{-1, -1}
}

// @lc code=end

