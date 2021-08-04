/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (40.54%)
 * Likes:    4306
 * Dislikes: 0
 * Total Accepted:    464.5K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,3], nums2 = [2]
 * 输出：2.00000
 * 解释：合并数组 = [1,2,3] ，中位数 2
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [1,2], nums2 = [3,4]
 * 输出：2.50000
 * 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums1 = [0,0], nums2 = [0,0]
 * 输出：0.00000
 *
 *
 * 示例 4：
 *
 *
 * 输入：nums1 = [], nums2 = [1]
 * 输出：1.00000
 *
 *
 * 示例 5：
 *
 *
 * 输入：nums1 = [2], nums2 = []
 * 输出：2.00000
 *
 *
 *
 *
 * 提示：
 *
 *
 * nums1.length == m
 * nums2.length == n
 * 0
 * 0
 * 1
 * -10^6
 *
 *
 *
 *
 * 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？
 *
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if l := (len(nums1) + len(nums2)); l%2 == 0 {
		//总个数为偶数时,如【1,2】需获取index 0 和 1 对应的值，这里先未判断数组是否为空，因为findkth里会判断
		return (findkth(nums1, nums2, l/2-1) + findkth(nums1, nums2, l/2)) / 2.0
	} else {
		//总个数为奇数时
		return findkth(nums1, nums2, l/2)
	}

}
func findkth(nums1 []int, nums2 []int, k int) float64 {
	for {
		l1, l2 := len(nums1), len(nums2)
		m1, m2 := l1/2, l2/2
		if l1 == 0 {
			return float64(nums2[k])
		} else if l2 == 0 {
			return float64(nums1[k])
		} else if k == 0 {
			if nums2[0] <= nums1[0] {
				return float64(nums2[0])
			} else {
				return float64(nums1[0])
			}
		}
		// k+1个刚好等于截取的这些m1+m2+2时
		if k == m1+m2+1 {
			//若相等则直接得到
			if nums1[m1] == nums2[m2] {
				return float64(nums1[m1])
				//否则，小块肯定不包括目标值，去掉小块
			} else if nums1[m1] < nums2[m2] {
				nums1 = nums1[m1+1:]
				k = k - m1 - 1
			} else {
				nums2 = nums2[m2+1:]
				k = k - m2 - 1
			}
			//如果小于，则小块可能包含目标值，大块肯定不包含，去掉大块
		} else if k < m1+m2+1 {
			if nums1[m1] <= nums2[m2] {
				nums2 = nums2[:m2]
			} else {
				nums1 = nums1[:m1]
			}
			//如果大于，则小块不包含目标值，去掉小块
		} else {
			if nums1[m1] <= nums2[m2] {
				nums1 = nums1[m1+1:]
				k = k - m1 - 1
			} else {
				nums2 = nums2[m2+1:]
				k = k - m2 - 1
			}
		}

	}
}

// @lc code=end

